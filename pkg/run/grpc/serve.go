package grpc

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grin-ch/grin-auth/cfg"
	etcd "github.com/grin-ch/grin-etcd-center"
	"github.com/grin-ch/grin-utils/log"
	"github.com/sirupsen/logrus"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

// RunServer 运行服务
func RunServer(serverName string, grpcPort int, r Registrar, connectors ...ClientConnector) error {
	return grpcServer(serverName, grpcPort, r, connectors...)
}

// 运行grpc服务
func grpcServer(serverName string, grpcPort int, r Registrar, connectors ...ClientConnector) error {
	// grpc listener
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", grpcPort))
	if err != nil {
		log.Logger.Errorf("tcp listen err:%s", err.Error())
		return err
	}

	opts := []grpc_logrus.Option{
		grpc_logrus.WithTimestampFormat("2006-01-02 15:04:05.000"),
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.cost", duration.Nanoseconds()
		}),
	}

	// grpc server
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryFunc)),
			grpc_auth.UnaryServerInterceptor(authFunc),
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(log.Logger), opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryFunc)),
			grpc_auth.StreamServerInterceptor(authFunc),
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(log.Logger), opts...),
		),
	)
	gracefulShutdown(s)

	// 获取注册中心连接
	etcdCenter, err := etcd.NewEtcdCenter(cfg.Config.Etcd.Endpoints, cfg.Config.Etcd.Timeout)
	if err != nil {
		log.Logger.Errorf("new etcd client err:%s", err.Error())
		return err
	}
	// 服务注册
	registrar := etcdCenter.Registrar(serverName, cfg.Config.Server.Host, grpcPort,
		etcd.WithLogger(func(err error) {
			log.Logger.Errorf("etcd err:%v", err)
		}))
	if err = registrar.Registry(); err != nil {
		log.Logger.Errorf("%s registry err: %s", serverName, err.Error())
		return err
	}
	defer func() {
		if err = registrar.Deregistry(); err != nil {
			log.Logger.Errorf("%s deregistry err: %s", serverName, err.Error())
		}
	}()
	// 初始化grpc客户端
	initClients(etcdResolver(etcdCenter.Builder()), connectors...)
	// 注册grpc服务
	registryServer(s, r)

	log.Logger.Infof("%s is running: %s", serverName,
		fmt.Sprintf("%s:%d", cfg.Config.Server.Host, grpcPort))
	return s.Serve(grpcListener)
}

func etcdResolver(builder resolver.Builder) func(string) (*grpc.ClientConn, error) {
	return func(s string) (*grpc.ClientConn, error) {
		addr := fmt.Sprintf("etcd:///%s", s)
		return grpc.Dial(addr,
			grpc.WithResolvers(builder),
			grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy":"%s"}`, roundrobin.Name)),
			grpc.WithKeepaliveParams(
				keepalive.ClientParameters{
					Time:                10 * time.Second,
					Timeout:             100 * time.Millisecond,
					PermitWithoutStream: true},
			),
			grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

// 优雅退出
func gracefulShutdown(svc *grpc.Server) {
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sign := <-quit

		go func() {
			time.Sleep(time.Duration(cfg.Config.Server.ForceStop) * time.Second)
			log.Logger.Errorf("grpc server force stop: sign:%v", sign)
			os.Exit(-1)
		}()
		// 关闭服务链接
		svc.GracefulStop()
		log.Logger.Infof("grpc server shutdown: %v", sign)
	}()
}

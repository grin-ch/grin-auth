package grpc

import (
	"fmt"
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"

	"github.com/grin-ch/grin-auth/cfg"
	etcd "github.com/grin-ch/grin-etcd-center"
	"github.com/grin-ch/grin-utils/log"
)

// RunServer 运行服务
func RunServer(serverName string, r Registrar, connectors ...ClientConnector) error {
	return grpcServer(serverName, r, connectors...)
}

// 运行grpc服务
func grpcServer(serverName string, r Registrar, connectors ...ClientConnector) error {
	// grpc listener
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Config.Server.Grpc.Port))
	if err != nil {
		log.Logger.Errorf("tcp listen err:%s", err.Error())
		return err
	}

	// grpc server
	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryFunc)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryFunc)),
		),
	)
	gracefulShutdown(s)
	// 注册grpc服务
	registryServer(s, r)

	// 获取注册中心连接
	etcdCenter, err := etcd.NewEtcdCenter(cfg.Config.Etcd.Endpoints, cfg.Config.Etcd.Timeout)
	if err != nil {
		log.Logger.Errorf("new etcd client err:%s", err.Error())
		return err
	}
	// 服务注册
	registrar := etcdCenter.Registrar(serverName, cfg.Config.Server.Host, cfg.Config.Server.Grpc.Port,
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

	log.Logger.Infof("%s is running: %s", serverName,
		fmt.Sprintf("%s:%d", cfg.Config.Server.Host, cfg.Config.Server.Grpc.Port))
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

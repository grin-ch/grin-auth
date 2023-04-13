package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-auth/pkg/util"
	"github.com/grin-ch/grin-utils/log"
)

func RunServer() {
	initClients()

	scfg := cfg.Config.Server.HttpServer
	r := Router(scfg.Mode)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", scfg.Info.Port),
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Errorf("server listen err:%s", err)
		}
	}()
	util.RunPprof(scfg.Info.PprofEnable, scfg.Info.PprofPort)
	log.Logger.Infof("%s is running: %s", scfg.Info.Name, fmt.Sprintf("%s:%d", cfg.Config.Server.Host, scfg.Info.Port))

	err := httpGracefulStop(
		server,
		time.Duration(scfg.Timeout),
		time.Duration(cfg.Config.Server.ForceStop),
	)

	if err != nil {
		log.Logger.Errorf("GracefulStop err:%s", err)
	}
}

func httpGracefulStop(server http.Server, timeout, forceStop time.Duration) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, channel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer channel()
	go func() {
		time.Sleep(forceStop * time.Second)
		panic("UnGracefulStop")
	}()
	return server.Shutdown(ctx)
}

func initClients() {
	// 获取注册中心连接
	etcdCenter, err := grpc.CreateEtcdCenter(cfg.Config.Etcd.Endpoints, cfg.Config.Etcd.Timeout)
	if err != nil {
		log.Logger.Errorf("new etcd client err:%s", err.Error())
		return
	}
	grpc.InitClients(etcdCenter, grpc.AllConn()...)
}

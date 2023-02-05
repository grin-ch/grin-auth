package grpc

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-utils/log"
	"google.golang.org/grpc"
)

type Registrar func(svc *grpc.Server)

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

// 注册grpc服务到本地主机
func registryServer(svc *grpc.Server, registry Registrar) {
	registry(svc)
}

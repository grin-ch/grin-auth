package grpc

import (
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/middle_ware/redis"
	service "github.com/grin-ch/grin-auth/pkg/service/captcha"

	"google.golang.org/grpc"
)

type Registrar func(svc *grpc.Server)

// 注册grpc服务到本地主机
func registryServer(svc *grpc.Server, registry Registrar) {
	registry(svc)
}

func RegistryCaptchaServer() Registrar {
	return func(svc *grpc.Server) {
		scfg := cfg.Config.Server.CaptchaServer
		captcha.RegisterCaptchaServiceServer(svc,
			service.NewCaptchaServer(redis.Client,
				scfg.Expires, scfg.Port, scfg.Subject, scfg.Format, scfg.Host, scfg.From, scfg.Secret))
	}
}

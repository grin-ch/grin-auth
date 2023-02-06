package grpc

import (
	"github.com/grin-ch/grin-api/api/grpc/account"
	"github.com/grin-ch/grin-api/api/grpc/captcha"
	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/middle_ware/mysql"
	"github.com/grin-ch/grin-auth/pkg/middle_ware/redis"
	account_impl "github.com/grin-ch/grin-auth/pkg/service/account"
	captcha_impl "github.com/grin-ch/grin-auth/pkg/service/captcha"

	"google.golang.org/grpc"
)

type Registrar func(svc *grpc.Server)

// 注册grpc服务到本地主机
func registryServer(svc *grpc.Server, registry Registrar) {
	registry(svc)
}

func RegistryCaptchaServices() Registrar {
	return func(svc *grpc.Server) {
		scfg := cfg.Config.Server.CaptchaServer
		captcha.RegisterCaptchaServiceServer(svc,
			captcha_impl.NewCaptchaServer(redis.Client,
				scfg.Expires, scfg.Port, scfg.Subject, scfg.Format, scfg.Host, scfg.From, scfg.Secret))
	}
}

func RegistryAccountServices() Registrar {
	return func(svc *grpc.Server) {
		scfg := cfg.Config.Server.AccountServer
		account.RegisterUserServiceServer(svc,
			account_impl.NewUserService(mysql.Client, scfg.Expires, scfg.Signed, scfg.Issuer, captchaClient))
	}
}

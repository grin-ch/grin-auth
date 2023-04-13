package main

import (
	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/middle_ware"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-utils/log"
)

func main() {
	cfg.InitConfig()
	scfg := cfg.Config.Server.CaptchaServer
	middle_ware.InitMiddleWare(
		middle_ware.InitRedis(),
	)

	log.InitLogger(
		log.WithLevel(cfg.Config.Log.Level),
		log.WithColor(cfg.Config.Log.Color),
		log.WithCaller(cfg.Config.Log.Caller),
		log.WithPath(scfg.Info.LogPath),
	)

	grpc.RunServer(
		scfg.Info,
		grpc.RegistryCaptchaServices(),
	)
}

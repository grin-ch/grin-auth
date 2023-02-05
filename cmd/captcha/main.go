package main

import (
	"math/rand"
	"time"

	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/middle_ware"
	"github.com/grin-ch/grin-auth/pkg/run/grpc"
	"github.com/grin-ch/grin-utils/log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg.InitConfig()
	scfg := cfg.Config.Server.CaptchaServer
	middle_ware.InitMiddleWare(
		middle_ware.InitRedis(),
	)

	log.InitLogger(
		log.WithLevel(cfg.Config.Log.Level),
		log.WithColor(cfg.Config.Log.Color),
		log.WithCaller(cfg.Config.Log.Caller),
		log.WithPath(scfg.LogPath),
	)

	grpc.RunServer(
		scfg.Name,
		grpc.RegistryCaptchaServer(),
	)
}

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
	scfg := cfg.Config.Server.AccountServer
	middle_ware.InitMiddleWare(
		middle_ware.InitMysql(cfg.Config.Dsn()),
	)

	log.InitLogger(
		log.WithLevel(cfg.Config.Log.Level),
		log.WithColor(cfg.Config.Log.Color),
		log.WithCaller(cfg.Config.Log.Caller),
		log.WithPath(scfg.Info.LogPath),
	)

	grpc.RunServer(
		scfg.Info.Name,
		scfg.Info.GrpcPort,
		grpc.RegistryAccountServices(),
		grpc.CaptchaConn,
	)
}

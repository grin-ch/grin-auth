package main

import (
	"github.com/grin-ch/grin-auth/cfg"
	grin_http "github.com/grin-ch/grin-auth/pkg/run/http"
	"github.com/grin-ch/grin-utils/log"
)

func main() {
	cfg.InitConfig()
	scfg := cfg.Config.Server.HttpServer

	log.InitLogger(
		log.WithLevel(cfg.Config.Log.Level),
		log.WithColor(cfg.Config.Log.Color),
		log.WithCaller(cfg.Config.Log.Caller),
		log.WithPath(scfg.Info.LogPath),
	)

	grin_http.RunServer()
}

package middle_ware

import (
	"github.com/grin-ch/grin-auth/cfg"
	"github.com/grin-ch/grin-auth/pkg/middle_ware/redis"
)

type initFunc func() error

func InitMiddleWare(initFuncs ...initFunc) {
	for _, fn := range initFuncs {
		if err := fn(); err != nil {
			panic(err)
		}
	}
}

func InitRedis() initFunc {
	return func() error {
		return redis.RedisInit(cfg.Config.Redis.Addr, cfg.Config.Redis.Pass, cfg.Config.Redis.DB)
	}
}

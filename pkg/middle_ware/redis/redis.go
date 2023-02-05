package redis

import "github.com/go-redis/redis"

var Client *redis.Client

func RedisInit(addr, pass string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	_, err := Client.Ping().Result()
	return err
}

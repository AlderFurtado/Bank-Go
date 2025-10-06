package cache

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	client *redis.Client
}

var instance *RedisCache
var once sync.Once

func GetRedisCache() *RedisCache {
	once.Do(func() {
		//TODO botar config em outro lugar
		rdb := redis.NewClient(&redis.Options{
			Addr:     "redis:6379", // "localhost:6379" ou "redis:6379" se estiver em docker
			Password: "",           // senha, se tiver
			DB:       0,            // banco padrão
		})

		instance = &RedisCache{client: rdb}
	})
	return instance
}

func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

package cache

import (
	"context"
	"time"

	"github.com/TheDevExperiment/server/internal/log"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const version = "v1" //change this to invalidate cache

type redisCache struct {
	client *redis.Client
}

func (rc *redisCache) start() {
	rdb := redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.addr"),
	})
	rc.client = rdb
}

func (rc *redisCache) Get(ctx context.Context, key string) (interface{}, error) {
	key = getCacheKey(key)
	val, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rc *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) {
	key = getCacheKey(key)
	err := rc.client.Set(ctx, key, value, ttl).Err()
	// TODO: implement compression of data before storing it
	if err != nil {
		log.Error(err)
	}
}

func getCacheKey(key string) string {
	return version + "_" + key
}

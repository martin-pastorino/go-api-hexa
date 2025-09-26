package db

import (
	"api/infra/config"
	"context"
	"fmt"
	"time"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value interface{}, ttlSeconds int) error
	Get(key string) (string, error)
	Delete(key string) error
}

func (lc *LocalCache) Set(key string, value interface{}, ttlSeconds int) error {
	ctx := context.Background()
	return lc.cache.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (lc *LocalCache) Get(key string) (string, error) {
	ctx := context.Background()
	return lc.cache.Get(ctx, key).Result()
}

func (lc *LocalCache) Delete(key string) error {
	ctx := context.Background()
	return lc.cache.Del(ctx, key).Err()
}

type LocalCache struct {
	cache *redis.Client
}

func NewCache(config *config.Config) *LocalCache {

	return &LocalCache{
		cache: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort),
			Password: config.RedisPassword, // no password set
			Username: "default",            // use default username
			DB:       0,                    // use default DB
		}),
	}
}

// Provider for UserRepository
func NewCacheProvider(config *config.Config) *LocalCache {
	return NewCache(config)
}

var ProviderSet = wire.NewSet(NewCache, wire.Bind(new(Cache), new(*LocalCache)))
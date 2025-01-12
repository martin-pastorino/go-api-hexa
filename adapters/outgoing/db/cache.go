package db

import (
	"api/infra/config"
	"fmt"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value interface{}, ttlSeconds int) error
	Get(key string) (interface{}, error)
	Delete(key string) error
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

var ProviderSet = wire.NewSet(NewCache,  wire.Bind(new(Cache), new(*LocalCache)))
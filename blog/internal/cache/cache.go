package cache

import (
	"blog/internal/cache/redis"

	"github.com/google/wire"
)

// ProviderSet is cache providers.
var ProviderSet = wire.NewSet(redis.NewRedisClient)

type Cache interface {
	Get(key string) (value string, err error)
	Set(key string, value string, ttl int) error
	HashGet(key string, fields *[]string) (values *map[string]string, err error)
	HashSet(key string, hash *map[string]string) error
}

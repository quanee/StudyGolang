package redis

import (
	"context"
	"sync"
	"time"

	"blog/internal/conf"

	"github.com/go-redis/redis/v8"
)

type RedisC struct {
	c       *redis.Client
	l       sync.Mutex
	address string
	size    int
}

func NewRedisClient(conf *conf.Cache) (*RedisC, error) {
	rc := &RedisC{
		address: conf.Connect.Source,
		size:    int(conf.Connect.Size),
	}
	if err := rc.GetRedisClient(); err != nil {
		return rc, err
	}
	return rc, nil
}

func (r *RedisC) GetRedisClient() error {
	if r.c == nil {
		r.l.Lock()
		defer r.l.Unlock()
		if r.c == nil {
			r.c = redis.NewClient(&redis.Options{
				Addr: r.address,
				Password: "",
				DB: 0,
			})
		}
	}
	return nil
}

func (r *RedisC) Get(key string) (string, error) {
	ctx := context.Background()
	ret, err := r.c.Get(ctx, key).Result()
	return ret, err
}

func (r *RedisC) Set(key string, value string, ttl int) error {
	ctx := context.Background()
	_, err := r.c.Set(ctx, key, value, time.Duration(ttl)).Result()
	return err
}

func (r *RedisC) HashGet(key string, fields *[]string) (values *map[string]string, err error) {
	ctx := context.Background()
	for _, field := range *fields {
		ret, err := r.c.HGet(ctx, key, field).Result()
		if err != nil {
			continue
		}
		(*values)[field] = ret
	}

	return values, err
}

func (r *RedisC) HashSet(key string, hash *map[string]string) error {
	ctx := context.Background()
	for k, v := range *hash {
		_, err := r.c.HMSet(ctx, key, k, v).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

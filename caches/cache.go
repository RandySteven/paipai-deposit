package caches

import (
	"context"

	redis_client "github.com/RandySteven/paipai-deposit/pkg/redis"
	"github.com/redis/go-redis/v9"
)

type (
	Caches interface {
		Set(ctx context.Context, key string, value any) error
		Get(ctx context.Context, key string) (any, error)
		GetMultiple(ctx context.Context, keys string) ([]any, error)
		SetMultiple(ctx context.Context, keys string, values []any) error
		Del(ctx context.Context, key string) error
	}

	caches struct {
		redis *redis.Client
	}
)

// Caches is a non-generic interface (values are untyped via any) so call sites can use
// caches.Caches without type arguments. A generic Caches[T] would require Caches[any] here.
var _ Caches = (*caches)(nil)

// Del implements [Caches].
func (c *caches) Del(ctx context.Context, key string) error {
	return redis_client.Del(ctx, c.redis, key)
}

// Get implements [Caches].
func (c *caches) Get(ctx context.Context, key string) (any, error) {
	return redis_client.Get(ctx, c.redis, key)
}

// Set implements [Caches].
func (c *caches) Set(ctx context.Context, key string, value any) error {
	return redis_client.Set(ctx, c.redis, key, value)
}

func (c *caches) GetMultiple(ctx context.Context, keys string) ([]any, error) {
	return redis_client.GetMultiple(ctx, c.redis, keys)
}

func (c *caches) SetMultiple(ctx context.Context, keys string, values []any) error {
	return redis_client.SetMultiple(ctx, c.redis, keys, values)
}

func NewCaches(redis *redis.Client) Caches {
	return &caches{
		redis: redis,
	}
}

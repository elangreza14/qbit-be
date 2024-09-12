package redislib

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
	ShutdownSave(ctx context.Context) *redis.StatusCmd
	Close() error
}

func NewRedis(ctx context.Context, redisAddr string, redisPassword string, redisDB int) (IRedis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

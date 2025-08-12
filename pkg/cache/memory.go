package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type memory struct {
	ctx context.Context
}

func (m memory) Get(key string, defaultValue func() (string, error)) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m memory) Set(key string, value interface{}, expiration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (m memory) Delete(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m memory) Pull(key string, defaultValue func() (string, error)) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m memory) Remember(key string, defaultValue func() (string, time.Duration, error)) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m memory) Close() error {
	//TODO implement me
	panic("implement me")
}

func New() Cache {
	return &memory{
		ctx: context.Background(),
	}
}

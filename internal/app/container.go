package app

import (
	redisDriver "github.com/go-redis/redis/v9"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/zap"
	"github.com/khodemobin/golang_boilerplate/pkg/pgsql"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"

	"gorm.io/gorm"
)

type Container struct {
	Cache  cache.Cache
	DB     *gorm.DB
	Redis  *redisDriver.Client
	Log    logger.Logger
	Config *config.Config
}

var c *Container = nil

func New() {
	cfg := config.New()

	clog := zap.New()

	db, err := pgsql.New(cfg)
	if err != nil {
		clog.Fatal(err)
	}

	rc := redis.New(cfg)
	if err != nil {
		clog.Fatal(err)
	}

	ca := cache.New(rc)

	c = &Container{
		Config: cfg,
		Log:    clog,
		DB:     db.DB,
		Cache:  ca,
	}
}

func App() *Container {
	return c
}

func Cache() cache.Cache {
	return c.Cache
}

func DB() *gorm.DB {
	return c.DB
}

func Log() logger.Logger {
	return c.Log
}

func Config() *config.Config {
	return c.Config
}

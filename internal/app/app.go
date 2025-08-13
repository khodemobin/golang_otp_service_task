package app

import (
	"github.com/khodemobin/golang_boilerplate/internal/model"
	"github.com/khodemobin/golang_boilerplate/internal/server/handler"
	"github.com/khodemobin/golang_boilerplate/pkg/logger/zap"
	"github.com/khodemobin/golang_boilerplate/pkg/pgsql"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/internal/service"
	"github.com/khodemobin/golang_boilerplate/pkg/cache"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"

	"gorm.io/gorm"
)

type App struct {
	Cache   cache.Cache
	DB      *gorm.DB
	Log     logger.Logger
	Config  *config.Config
	Service *service.Service
	handler *handler.Handler
}

func New() *App {
	cfg := config.New()

	logVar := zap.New()

	db, err := pgsql.New(cfg)
	if err != nil {
		logVar.Fatal(err)
	}

	err = db.DB.AutoMigrate(model.User{})
	if err != nil {
		logVar.Fatal(err)
	}

	cacheVar := cache.New()

	return &App{
		Config:  cfg,
		Log:     logVar,
		DB:      db.DB,
		Cache:   cacheVar,
		Service: service.NewService(cfg, db.DB, logVar, cacheVar),
	}
}

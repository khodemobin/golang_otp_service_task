package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Port      int    `env:"APP_PORT" env-default:"8000"`
	Env       string `env:"APP_ENV" env-default:"local"`
	JWTSecret string `env:"JWT_SECRET" env-default:"secret"`
}

type DB struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Database string `env:"DB_DATABASE" env-default:"test"`
	User     string `env:"DB_USERNAME" env-default:"test"`
	Password string `env:"DB_PASSWORD" env-default:"secret"`
	Port     int    `env:"DB_PORT" env-default:"3306"`
}

type Config struct {
	App App
	DB  DB
}

var cfg Config

func New() *Config {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(b), "../..")

	err := cleanenv.ReadConfig(path+"/.env", &cfg)
	if err != nil {
		log.Fatalln("error load config", err)
	}

	return &cfg
}

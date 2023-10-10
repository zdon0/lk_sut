package config

import (
	"github.com/caarlos0/env/v9"
	"time"
)

type Config struct {
	Api          Api
	LkSutService LkSutService
	Scheduler    Scheduler
	Redis        Redis
	Debug        bool `env:"APP_DEBUG" envDefault:"true"`
}

type Api struct {
	Addr              string        `env:"API_HOST" envDefault:"0.0.0.0"`
	Port              int           `env:"API_PORT" envDefault:"8080"`
	ReadHeaderTimeout time.Duration `env:"API_TIMEOUT" envDefault:"10s"`
}

type LkSutService struct {
	URL     string        `env:"LK_SUT_URL" envDefault:"https://lk.sut.ru"`
	Timeout time.Duration `env:"LK_SUT_TIMEOUT," envDefault:"10s"`
}

type Scheduler struct {
	CommitterInterval time.Duration `env:"SCHEDULER_COMMIT_INTERVAL" envDefault:"3m"`
}

type Redis struct {
	Addr                string        `env:"REDIS_HOST,required"`
	Port                int           `env:"REDIS_PORT" envDefault:"6379"`
	Password            string        `env:"REDIS_PASSWORD,required"`
	DB                  int           `env:"REDIS_DB" envDefault:"0"`
	Timeout             time.Duration `env:"REDIS_TIMEOUT" envDefault:"10s"`
	UserDataHTable      string        `env:"REDIS_USER_DATA_TABLE" envDefault:"lk-sut:user:data"`
	UserLastLoginHTable string        `env:"REDIS_USER_LOGIN_TABLE" envDefault:"lk-sut:user:login"`
}

func NewConfig() *Config {
	var res Config

	if err := env.Parse(&res); err != nil {
		panic(err)
	}

	return &res
}

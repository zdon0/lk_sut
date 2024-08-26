package testutils

import (
	"net/http/httptest"
	"testing"
	_ "time/tzdata" // for MSK

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redismock/v9"
	"github.com/maxcnunes/httpfake"
	"go.uber.org/fx"
	"go.uber.org/zap/zaptest"

	"lk_sut/internal/config"
	"lk_sut/internal/di"
	"lk_sut/internal/utils"
	"lk_sut/internal/worker"
)

type App struct {
	Config *config.Config
	Api    *httptest.Server

	sutServer *httpfake.HTTPFake
	redisMock redismock.ClientMock

	ClearFunc func()
}

type Opt func(*App)

func NewMockApi(t testing.TB, opts ...Opt) *App {
	return newMockApp(t, false, opts...)
}

func NewMockApiWithWorker(t testing.TB, opts ...Opt) *App {
	return newMockApp(t, true, opts...)
}

func newMockApp(t testing.TB, needWorker bool, opts ...Opt) *App {
	cfg := NewConfig()

	redisClient, redisClientMock := redismock.NewClientMock()

	sutServer := httpfake.New(httpfake.WithTesting(t))
	cfg.LkSutService.URL = sutServer.Server.URL

	res := App{
		Config:    cfg,
		sutServer: sutServer,
		redisMock: redisClientMock,
	}

	for _, opt := range opts {
		opt(&res)
	}

	invoke := []any{
		utils.SetLocation,
		func(api *gin.Engine) {
			res.Api = httptest.NewServer(api)
		},
	}

	if needWorker {
		invoke = append(invoke, worker.InitializeWorker)
	}

	fxInstance := fx.New(
		fx.Options(
			fx.Provide(di.Constructors()...),
			fx.Replace(
				res.Config,
				zaptest.NewLogger(t),
				redisClient,
			),
			fx.Invoke(invoke...),
			fx.NopLogger,
		),
	)

	closeFx := startFxInstance(t, fxInstance)

	res.ClearFunc = func() {
		closeFx()

		if err := redisClientMock.ExpectationsWereMet(); err != nil {
			t.Errorf("redis mock expectations: %s", err)
		}

		res.Api.Close()
		res.sutServer.Close()
		_ = redisClient.Close()
	}

	return &res
}

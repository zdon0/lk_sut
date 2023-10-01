package testutils

import (
	"github.com/go-redis/redismock/v9"
	"github.com/maxcnunes/httpfake"
	"go.uber.org/zap/zaptest"
	"lk_sut/internal/api"
	userHandler "lk_sut/internal/api/handler/user"
	"lk_sut/internal/config"
	"lk_sut/internal/env"
	userInteractor "lk_sut/internal/interactor/user"
	"lk_sut/internal/repository/user"
	"lk_sut/internal/sutclient"
	"lk_sut/internal/worker"
	"net/http"
	"net/http/httptest"
	"testing"
)

type App struct {
	Config *config.Config
	Api    *httptest.Server

	SutServer *httpfake.HTTPFake
	RedisMock redismock.ClientMock
	Worker    *worker.Worker

	ClearFunc func()
}

func NewMockApp(t testing.TB) *App {
	var cfg config.Config

	logger := zaptest.NewLogger(t)

	redisClient, redisClientMock := redismock.NewClientMock()

	sutServer := httpfake.New(httpfake.WithTesting(t))
	cfg.LkSutService.URL = sutServer.Server.URL
	cfg.LkSutService.Timeout = 0

	cfg.Redis.UserDataHTable = RedisHashTableData
	cfg.Redis.UserLastLoginHTable = RedisHashTableLogin
	userRepo := user.NewRepo(redisClient, cfg.Redis)

	handlerLkSutClient := sutclient.NewClient(cfg.LkSutService)
	interactor := userInteractor.NewInteractor(userRepo, handlerLkSutClient)

	handler := userHandler.NewHandler(interactor)
	e := env.NewEnv(handler)

	server := api.NewApi(&cfg, e, logger)

	workerLkSutClient := sutclient.NewClient(cfg.LkSutService)
	lkSutClientWorker := worker.NewWorker(cfg.Scheduler, workerLkSutClient, userRepo, logger)

	res := App{
		Config:    &cfg,
		Api:       httptest.NewServer(server.Handler),
		SutServer: sutServer,
		RedisMock: redisClientMock,
		Worker:    lkSutClientWorker,
	}

	res.ClearFunc = func() {
		if err := redisClientMock.ExpectationsWereMet(); err != nil {
			t.Errorf("redis mock expectations: %s", err)
		}

		res.Api.Close()
		res.SutServer.Close()
		res.Worker.Stop()
		_ = redisClient.Close()
	}

	return &res
}

func (a *App) RegisterUserInitHandler() {
	a.SutServer.NewHandler().Get("/cabinet").
		AssertHeaderValue("Cookie", "").
		Reply(http.StatusOK).
		AddHeader(SetCookieHeader, UidCookieHeaderResponse())
}

func (a *App) RegisterUserAuthHandler(login string, password string) {
	a.SutServer.NewHandler().
		Post("/cabinet/lib/autentificationok.php").
		AssertCustom(newAuthAssertor(login, password)).
		AssertHeaderValue(CookieHeader, UidCookieHeaderRequest()).
		Reply(http.StatusOK).
		BodyString("1")
}

func (a *App) RegisterUserAuthHandlerBadUser(login, password string) {
	a.SutServer.NewHandler().
		Post("/cabinet/lib/autentificationok.php").
		AssertCustom(newAuthAssertor(login, password)).
		AssertHeaderValue(CookieHeader, UidCookieHeaderRequest()).
		Reply(http.StatusOK).
		BodyString("0")
}

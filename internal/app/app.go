package app

import (
	"errors"
	"lk_sut/internal/api"
	userHandler "lk_sut/internal/api/handler/user"
	"lk_sut/internal/config"
	"lk_sut/internal/db/redis"
	"lk_sut/internal/env"
	userInteractor "lk_sut/internal/interactor/user"
	appLogger "lk_sut/internal/logger"
	"lk_sut/internal/repository/user"
	"lk_sut/internal/sutclient"
	"lk_sut/internal/worker"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	time.Local = time.UTC

	cfg := config.NewConfig()

	logger, err := appLogger.NewLogger(cfg)
	if err != nil {
		return err
	}

	defer func() {
		_ = logger.Sync()
	}()

	db, err := redis.NewClient(cfg.Redis)
	if err != nil {
		return err
	}

	defer func() {
		_ = db.Close()
	}()

	userRepo := user.NewRepo(db, cfg.Redis)

	handlerLkSutClient := sutclient.NewClient(cfg.LkSutService)
	interactor := userInteractor.NewInteractor(userRepo, handlerLkSutClient)

	handler := userHandler.NewHandler(interactor)
	e := env.NewEnv(handler)

	server := api.NewApi(cfg, e, logger)

	workerLkSutClient := sutclient.NewClient(cfg.LkSutService)
	lkSutClientWorker := worker.NewWorker(cfg.Scheduler, workerLkSutClient, userRepo, logger)

	if err := lkSutClientWorker.RegisterJobs(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	lkSutClientWorker.StartAsync()

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	<-quit

	_ = server.Close()
	lkSutClientWorker.Stop()

	return <-errChan
}

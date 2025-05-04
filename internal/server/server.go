package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"lk_sut/internal/config"
)

func InitializeServer(cfg *config.Config, logger *zap.Logger, api *gin.Engine, shutdowner fx.Shutdowner, lc fx.Lifecycle) {
	srv := &http.Server{
		Addr:              net.JoinHostPort(cfg.Api.Addr, strconv.Itoa(cfg.Api.Port)),
		ReadHeaderTimeout: cfg.Api.ReadHeaderTimeout,
		Handler:           api,
	}

	r := runner{
		srv:        srv,
		logger:     logger,
		shutdowner: shutdowner,
	}

	lc.Append(fx.Hook{
		OnStart: r.Start,
		OnStop:  r.Stop,
	})
}

type runner struct {
	srv        *http.Server
	logger     *zap.Logger
	shutdowner fx.Shutdowner
}

func (r *runner) listenAndServe() {
	r.logger.Info("http server starting")

	err := r.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		r.logger.Error("http server stopped with error", zap.Error(err))
		_ = r.shutdowner.Shutdown()

		return
	}

	r.logger.Info("http server stopped")
}

func (r *runner) Start(_ context.Context) error {
	go r.listenAndServe()
	return nil
}

func (r *runner) Stop(ctx context.Context) error {
	return r.srv.Shutdown(ctx)
}

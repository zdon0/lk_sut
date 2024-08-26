package di

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"lk_sut/internal/api"
	userHandler "lk_sut/internal/api/handler/user"
	"lk_sut/internal/config"
	"lk_sut/internal/db/redis"
	domainUser "lk_sut/internal/domain/user"
	interactorUser "lk_sut/internal/interactor/user"
	"lk_sut/internal/logger"
	repoUser "lk_sut/internal/repository/user"
	"lk_sut/internal/server"
	"lk_sut/internal/sutclient"
	"lk_sut/internal/utils"
	"lk_sut/internal/worker"
)

func Constructors() []any {
	return []any{
		config.NewConfig,
		logger.NewLogger,

		redis.NewClient,

		fx.Annotate(
			repoUser.NewRepo,
			fx.As(new(domainUser.Repository)),
			fx.As(new(worker.Repository)),
		),

		fx.Annotate(sutclient.NewClient, fx.As(new(domainUser.Authorization))),
		fx.Annotate(sutclient.NewClient, fx.As(new(worker.LkSutCommitter))),

		domainUser.NewBehavior,

		interactorUser.NewInteractor,

		userHandler.NewHandler,

		api.NewApi,
	}
}

func CreateApp() fx.Option {
	return fx.Options(
		fx.Provide(Constructors()...),
		fx.Invoke(
			utils.SetLocation,
			server.InitializeServer,
			worker.InitializeWorker,
		),
		fx.WithLogger(func() fxevent.Logger {
			return logger.StartStopLogger{}
		}),
	)
}

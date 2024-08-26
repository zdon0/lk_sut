package worker

import (
	"context"
	"time"

	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"lk_sut/internal/config"
	"lk_sut/internal/pkg/daystamp"
)

type worker struct {
	scheduler         gocron.Scheduler
	committerInterval time.Duration

	lessonStartList []daystamp.Stamp
	lessonDuration  time.Duration

	logger *zap.Logger

	lkSut LkSutCommitter
	repo  Repository
}

func InitializeWorker(cfg *config.Config, lkSut LkSutCommitter, repo Repository, logger *zap.Logger, lc fx.Lifecycle) error {
	scheduler, err := gocron.NewScheduler(
		gocron.WithGlobalJobOptions(
			gocron.WithSingletonMode(gocron.LimitModeWait),
		),
	)
	if err != nil {
		return err
	}

	w := worker{
		scheduler:         scheduler,
		committerInterval: cfg.Scheduler.CommitterInterval,
		lessonStartList:   cfg.Lesson.StartList,
		lessonDuration:    cfg.Lesson.Duration,
		logger:            logger,
		lkSut:             lkSut,
		repo:              repo,
	}

	if err := w.registerJobs(); err != nil {
		return err
	}

	lc.Append(fx.Hook{
		OnStart: w.Start,
		OnStop:  w.Stop,
	})

	return nil
}

func (w *worker) Start(_ context.Context) error {
	w.logger.Info(workerMsg,
		zap.String(actionField, "worker started"))

	w.scheduler.Start()

	return nil
}

func (w *worker) Stop(_ context.Context) error {
	err := w.scheduler.Shutdown()

	w.logger.Info(workerMsg,
		zap.String(actionField, "worker stopped"))

	return err
}

func (w *worker) registerJobs() error {
	if err := w.registerCommitterJob(); err != nil {
		return err
	}

	return w.registerFlusherJob()
}

func (w *worker) registerCommitterJob() error {
	_, err := w.scheduler.NewJob(
		gocron.DurationJob(w.committerInterval),
		gocron.NewTask(w.commit),
	)

	return err
}

func (w *worker) registerFlusherJob() error {
	_, err := w.scheduler.NewJob(
		gocron.DailyJob(1,
			gocron.NewAtTimes(
				gocron.NewAtTime(0, 0, 0),
			),
		),
		gocron.NewTask(w.flush),
	)

	return err
}

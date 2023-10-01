package worker

import (
	"context"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"lk_sut/internal/config"
	"time"
)

type Worker struct {
	scheduler         *gocron.Scheduler
	schedulerInterval time.Duration

	logger *zap.Logger

	lkSut LkSutCommitter
	repo  Repository

	baseCtx context.Context
	cancel  func()
}

func NewWorker(cfg config.Scheduler, lkSut LkSutCommitter, repo Repository, logger *zap.Logger) *Worker {
	ctx, cancel := context.WithCancel(context.Background())

	return &Worker{
		scheduler:         gocron.NewScheduler(time.UTC),
		schedulerInterval: cfg.RepeatInterval,
		logger:            logger,
		lkSut:             lkSut,
		repo:              repo,
		baseCtx:           ctx,
		cancel:            cancel,
	}
}

func (w *Worker) Start() {
	w.logger.Info(workerMsg,
		zap.String(actionField, "worker started"))

	w.scheduler.StartBlocking()
}

func (w *Worker) StartAsync() {
	w.logger.Info(workerMsg,
		zap.String(actionField, "worker started"))

	w.scheduler.StartAsync()
}

func (w *Worker) Stop() {
	w.cancel()
	w.scheduler.Stop()

	w.logger.Info(workerMsg,
		zap.String(actionField, "worker stopped"))
}

func (w *Worker) RegisterJob() error {
	_, err := w.scheduler.
		Every(w.schedulerInterval).
		SingletonMode().
		Do(comitterFunc(w.lkSut, w.repo, w.logger), w.baseCtx)

	return err
}

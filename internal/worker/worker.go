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
	committerInterval time.Duration

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
		committerInterval: cfg.CommitterInterval,
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

func (w *Worker) RegisterJobs() error {
	if err := w.registerCommitterJob(); err != nil {
		return err
	}

	return w.registerFlusherJob()
}

func (w *Worker) registerCommitterJob() error {
	_, err := w.scheduler.
		Every(w.committerInterval).
		SingletonMode().
		Do(comitterFunc(w.lkSut, w.repo, w.logger), w.baseCtx)

	return err
}

func (w *Worker) registerFlusherJob() error {
	_, err := w.scheduler.
		Every(1).Day().At("00:00").
		SingletonMode().
		Do(flusherFunc(w.repo, w.logger), w.baseCtx)

	return err
}

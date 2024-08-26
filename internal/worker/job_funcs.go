package worker

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"lk_sut/internal/domain"
	"lk_sut/internal/pkg/daystamp"
	"lk_sut/internal/sutclient"
)

func (w *worker) commit() {
	curTime := time.Now()

	var (
		currentLesson daystamp.Stamp
		timeToCommit  bool
	)

	for _, lesson := range w.lessonStartList {
		if lesson.IsTimeInStamp(curTime, w.lessonDuration) {
			currentLesson = lesson
			timeToCommit = true

			break
		}
	}

	if !timeToCommit {
		return
	}

	ctx := context.TODO()

	list, err := w.repo.GetAllUsers(ctx)
	if err != nil {
		w.logger.Error(workerMsg,
			zap.String(actionField, "get all users"),
			zap.Error(err))

		return
	}

	for _, user := range list {
		lastLogin, err := w.repo.GetUserLastLogin(ctx, user.Login)
		if err != nil && !errors.Is(err, domain.ErrNotFound) {
			w.logger.Error(workerMsg,
				zap.String(loginField, user.Login),
				zap.String(actionField, "get last login"),
				zap.Error(err))

			continue
		}

		if err == nil && currentLesson.IsTimeInStamp(lastLogin, w.lessonDuration) {
			continue
		}

		w.logger.Debug(workerMsg,
			zap.String(loginField, user.Login),
			zap.String(actionField, "start commit"))

		if err := w.lkSut.AuthorizeUser(ctx, user.Login, user.Password); err != nil {
			w.logger.Error(workerMsg,
				zap.String(loginField, user.Login),
				zap.String(actionField, "authorize"),
				zap.Error(err))

			continue
		}

		if err := w.lkSut.CommitLesson(ctx); err != nil {
			if !errors.Is(err, sutclient.ErrNoLessonToCommit) {
				w.logger.Error(workerMsg,
					zap.String(loginField, user.Login),
					zap.String(actionField, "commit lesson"),
					zap.Error(err))
			}

			continue
		}

		w.logger.Info(workerMsg,
			zap.String(loginField, user.Login),
			zap.String(actionField, "committed"))

		if err := w.repo.SetUserLastLogin(ctx, user.Login, curTime); err != nil {
			w.logger.Error(workerMsg,
				zap.String(loginField, user.Login),
				zap.String(actionField, "set last login"),
				zap.Error(err))
		}
	}
}

func (w *worker) flush() {
	ctx := context.TODO()

	if err := w.repo.FlushLastLogin(ctx); err != nil {
		w.logger.Error(workerMsg,
			zap.String(actionField, "flush last login table"),
			zap.Error(err))

		return
	}

	w.logger.Debug(workerMsg,
		zap.String(actionField, "flushed last login table"))
}

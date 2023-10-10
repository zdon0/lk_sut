package worker

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"lk_sut/internal/domain"
	lessonPkg "lk_sut/internal/pkg/lesson"
	"lk_sut/internal/sutclient"
	"time"
)

type LkSutCommitter interface {
	AuthorizeUser(ctx context.Context, user domain.User) error
	CommitLesson(ctx context.Context) error
}

type Repository interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserLastLogin(ctx context.Context, login string) (time.Time, error)
	SetUserLastLogin(ctx context.Context, login string, timeToSet time.Time) error
	FlushLastLogin(ctx context.Context) error
}

func comitterFunc(lkSut LkSutCommitter, repo Repository, logger *zap.Logger) func(context.Context) {
	availableLessons := domain.AvailableLessons()

	return func(ctx context.Context) {
		curTime := time.Now()

		var timeToCommit bool
		var currentLesson domain.Lesson

		for _, lesson := range availableLessons {
			if lessonPkg.IsTimeStampInLesson(curTime, lesson) {
				currentLesson = lesson
				timeToCommit = true
				break
			}
		}

		if !timeToCommit {
			return
		}

		list, err := repo.GetAllUsers(ctx)
		if err != nil {
			logger.Error(workerMsg,
				zap.String(actionField, "get all users"),
				zap.Error(err))

			return
		}

		for _, user := range list {
			lastLogin, err := repo.GetUserLastLogin(ctx, user.Login)
			if err != nil && !errors.Is(err, domain.ErrNotFound) {
				logger.Error(workerMsg,
					zap.String(loginField, user.Login),
					zap.String(actionField, "get last login"),
					zap.Error(err))

				continue
			}

			if lessonPkg.IsTimeStampInLesson(lastLogin, currentLesson) {
				continue
			}

			logger.Debug(workerMsg,
				zap.String(loginField, user.Login),
				zap.String(actionField, "start commit"))

			if err := lkSut.AuthorizeUser(ctx, user); err != nil {
				logger.Error(workerMsg,
					zap.String(loginField, user.Login),
					zap.String(actionField, "authorize"),
					zap.Error(err))

				continue
			}

			if err := lkSut.CommitLesson(ctx); err != nil {
				if !errors.Is(err, sutclient.ErrNoLessonToCommit) {
					logger.Error(workerMsg,
						zap.String(loginField, user.Login),
						zap.String(actionField, "commit lesson"),
						zap.Error(err))
				}

				continue
			}

			logger.Info(workerMsg,
				zap.String(loginField, user.Login),
				zap.String(actionField, "committed"))

			if err := repo.SetUserLastLogin(ctx, user.Login, curTime); err != nil {
				logger.Error(workerMsg,
					zap.String(loginField, user.Login),
					zap.String(actionField, "set last login"),
					zap.Error(err))
			}
		}
	}
}

func flusherFunc(repo Repository, logger *zap.Logger) func(context.Context) {
	return func(ctx context.Context) {
		if err := repo.FlushLastLogin(ctx); err != nil {
			logger.Error(workerMsg,
				zap.String(actionField, "flush last login table"),
				zap.Error(err))

			return
		}

		logger.Debug(workerMsg,
			zap.String(actionField, "flushed last login table"))
	}
}

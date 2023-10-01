package lesson_test

import (
	"github.com/stretchr/testify/assert"
	"lk_sut/internal/domain"
	"lk_sut/internal/pkg/lesson"
	"testing"
	"time"
)

func TestIsTimeStampInLesson(t *testing.T) {
	defaultLesson := domain.Lesson{
		LeftBorder: domain.TimeStamp{
			Hour:   13,
			Minute: 00,
		},
		RightBorder: domain.TimeStamp{
			Hour:   14,
			Minute: 35,
		},
	}

	testcases := []struct {
		Lesson    domain.Lesson
		TimeStamp time.Time
		IsLesson  bool
	}{
		{
			Lesson:    defaultLesson,
			TimeStamp: time.Date(2023, 2, 14, 11, 15, 0, 0, time.UTC), // 14:15 (UTC+3)
			IsLesson:  true,
		},
		{
			Lesson:    defaultLesson,
			TimeStamp: time.Date(2023, 2, 14, 9, 50, 0, 0, time.UTC), // 12:50 (UTC+3)
			IsLesson:  false,
		},
		{
			Lesson:    defaultLesson,
			TimeStamp: time.Date(2023, 2, 14, 12, 15, 0, 0, time.UTC), // 15:15 (UTC+3)
			IsLesson:  false,
		},
	}

	for _, test := range testcases {
		assert.Equal(t, test.IsLesson, lesson.IsTimeStampInLesson(test.TimeStamp, test.Lesson))
	}
}

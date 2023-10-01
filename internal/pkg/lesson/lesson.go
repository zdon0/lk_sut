package lesson

import (
	"lk_sut/internal/domain"
	"time"
)

const moscowUTC = 3

func IsTimeStampInLesson(t time.Time, lesson domain.Lesson) bool {
	t = t.UTC().Add(time.Hour * moscowUTC)

	if t.Weekday() == time.Sunday {
		return false
	}

	currentHour, currentMinute := t.Hour(), t.Minute()

	lessonLeftHour, lessonLeftMinute := lesson.LeftBorder.Hour, lesson.LeftBorder.Minute
	lessonRightHour, lessonRightMinute := lesson.RightBorder.Hour, lesson.RightBorder.Minute

	curTimeStamp := timeStampFromHourAndMinute(t, currentHour, currentMinute)
	leftTimeStamp := timeStampFromHourAndMinute(t, lessonLeftHour, lessonLeftMinute)
	rightTimeStamp := timeStampFromHourAndMinute(t, lessonRightHour, lessonRightMinute)

	return leftTimeStamp.Before(curTimeStamp) && curTimeStamp.Before(rightTimeStamp)
}

func timeStampFromHourAndMinute(current time.Time, hour, minute int) time.Time {
	return time.Date(current.Year(), current.Month(), current.Day(), hour, minute, 0, 0, time.UTC)
}

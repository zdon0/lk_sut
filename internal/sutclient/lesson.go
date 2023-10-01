package sutclient

import (
	"context"
	"errors"
	"lk_sut/internal/pkg/decoder"
	"net/url"
	"regexp"
	"strings"
)

const (
	noButtonOnCurrentLesson = `ждем начала от преподавателя`
)

var openLessonParamsRe = regexp.MustCompile(`onclick="open_zan\((\d+),(\d+)\);"`)
var lessonStartedRe = regexp.MustCompile(`<td align="left">.*\d{2}:\d{2}<\/span><\/td>`)

func (sc *SutClient) matchOpenLessonParams(schedule string) (lesson string, week string, err error) {
	match := openLessonParamsRe.FindStringSubmatch(schedule)

	if len(match) == 0 {
		return "", "", ErrNoLessonToCommit
	}

	if len(match) != 3 {
		return "", "", ErrBadLessonParams
	}

	return match[1], match[2], nil
}

func (sc *SutClient) getSchedule(ctx context.Context) (string, error) {
	resp, err := sc.r(ctx).Post("/cabinet/project/cabinet/forms/raspisanie.php")
	if err != nil {
		return "", err
	}

	schedule, err := decoder.Decode(resp.String())
	if err != nil {
		return "", err
	}

	return schedule, nil
}

func (sc *SutClient) openLesson(ctx context.Context, schedule string) error {
	lesson, week, err := sc.matchOpenLessonParams(schedule)
	if err != nil {
		if errors.Is(err, ErrNoLessonToCommit) {
			if lessonStartedRe.MatchString(schedule) {
				return nil
			}

			// No lesson
			if !strings.Contains(schedule, noButtonOnCurrentLesson) {
				return nil
			}

			return err
		}

		return err
	}

	formData := make(url.Values, 3)

	formData.Set("open", "1")
	formData.Set("rasp", lesson)
	formData.Set("week", week)

	_, err = sc.r(ctx).
		SetFormDataFromValues(formData).
		Post("/cabinet/project/cabinet/forms/raspisanie.php")

	return err
}

func (sc *SutClient) CommitLesson(ctx context.Context) error {
	schedule, err := sc.getSchedule(ctx)
	if err != nil {
		return err
	}

	return sc.openLesson(ctx, schedule)
}

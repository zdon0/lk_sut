package utils

import (
	"time"

	"lk_sut/internal/config"
)

func SetLocation(cfg *config.Config) error {
	loc, err := time.LoadLocation(cfg.Lesson.Timezone)
	if err != nil {
		return err
	}

	time.Local = loc

	return nil
}

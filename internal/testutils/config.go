package testutils

import (
	"time"

	"lk_sut/internal/config"
	"lk_sut/internal/pkg/daystamp"
)

func NewConfig() *config.Config {
	now := time.Now()

	return &config.Config{
		Api:          config.Api{},
		LkSutService: config.LkSutService{},
		Scheduler: config.Scheduler{
			CommitterInterval: WorkerCommitInterval,
		},
		Redis: config.Redis{
			UserDataHTable:      RedisHashTableData,
			UserLastLoginHTable: RedisHashTableLogin,
		},
		Lesson: config.Lesson{
			Timezone: "Europe/Moscow",
			Duration: time.Minute,
			StartList: []daystamp.Stamp{
				daystamp.NewStamp(now.Hour(), now.Minute()),
			},
		},
	}
}

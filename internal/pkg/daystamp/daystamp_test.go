package daystamp_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"lk_sut/internal/pkg/daystamp"
	"lk_sut/internal/testutils"
	"lk_sut/internal/utils"
)

func TestStamp_IsTimeInStamp(t *testing.T) {
	err := utils.SetLocation(testutils.NewConfig())
	if err != nil {
		t.Fatal(err)
	}

	defaultDayStamp := daystamp.NewStamp(13, 0)
	defaultDuration := time.Hour + 30*time.Minute

	testcases := []struct {
		timestamp time.Time
		res       bool
	}{
		{
			timestamp: time.Date(2023, 2, 14, 14, 15, 0, 0, time.Local),
			res:       true,
		},
		{
			timestamp: time.Date(2023, 2, 14, 12, 50, 0, 0, time.Local),
			res:       false,
		},
		{
			timestamp: time.Date(2023, 2, 14, 15, 15, 0, 0, time.Local),
			res:       false,
		},
	}

	for _, tt := range testcases {
		assert.Equal(t, tt.res, defaultDayStamp.IsTimeInStamp(tt.timestamp, defaultDuration))
	}
}

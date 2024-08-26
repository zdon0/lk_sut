package daystamp

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

func NewStamp(hour, minute int) Stamp {
	return Stamp{
		hour:   hour,
		minute: minute,
	}
}

type Stamp struct {
	hour   int
	minute int
}

func (s *Stamp) IsTimeInStamp(t time.Time, duration time.Duration) bool {
	ts := time.Date(t.Year(), t.Month(), t.Day(), s.hour, s.minute, 0, 0, time.Local)

	return ts.Before(t) && t.Before(ts.Add(duration))
}

func (s *Stamp) UnmarshalText(text []byte) error {
	_, err := fmt.Fscanf(bytes.NewBuffer(text), "%d:%d", &s.hour, &s.minute)
	if err != nil {
		return err
	}

	if s.hour < 0 || s.hour > 23 {
		return errors.New("bad hour")
	}

	if s.minute < 0 || s.minute > 59 {
		return errors.New("bad minute")
	}

	return nil
}

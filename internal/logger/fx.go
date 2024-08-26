package logger

import (
	"fmt"

	"go.uber.org/fx/fxevent"
)

var _ fxevent.Logger = StartStopLogger{}

type StartStopLogger struct{}

func (l StartStopLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Started:
		if e.Err != nil {
			println(fmt.Sprintf("start failed: %v", e.Err))
		}
	case *fxevent.Stopped:
		if e.Err != nil {
			println(fmt.Sprintf("stop failed: %v", e.Err))
		}
	}
}

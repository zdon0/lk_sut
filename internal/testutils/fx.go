package testutils

import (
	"context"
	"testing"
	"time"

	"go.uber.org/fx"
)

func startFxInstance(t testing.TB, fxInstance *fx.App) func() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	err := fxInstance.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}

	stopFn := func() {
		ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelCtx()

		err := fxInstance.Stop(ctx)
		if err != nil {
			t.Fatal(err)
		}
	}

	return stopFn
}

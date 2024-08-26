package di_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/fx"

	"lk_sut/internal/di"
)

func TestCreateApp(t *testing.T) {
	require.NoError(t, fx.ValidateApp(di.CreateApp()))
}

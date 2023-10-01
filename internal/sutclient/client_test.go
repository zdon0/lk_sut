package sutclient_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"lk_sut/internal/domain"
	"lk_sut/internal/sutclient"
	"lk_sut/internal/testutils"
	"testing"
)

func TestSutClient_SetUser(t *testing.T) {
	api := testutils.NewMockApp(t)

	t.Cleanup(api.ClearFunc)

	api.RegisterUserInitHandler()
	api.RegisterUserAuthHandler(testutils.Login, testutils.Password)

	client := sutclient.NewClient(api.Config.LkSutService)

	user := domain.User{
		Login:    testutils.Login,
		Password: testutils.Password,
	}

	err := client.AuthorizeUser(context.Background(), user)
	require.NoError(t, err)
}

func TestSutClient_SetUser_bad_user(t *testing.T) {
	api := testutils.NewMockApp(t)

	t.Cleanup(api.ClearFunc)

	api.RegisterUserInitHandler()
	api.RegisterUserAuthHandlerBadUser(testutils.Login, testutils.Password)

	client := sutclient.NewClient(api.Config.LkSutService)

	user := domain.User{
		Login:    testutils.Login,
		Password: testutils.Password,
	}

	err := client.AuthorizeUser(context.Background(), user)
	require.ErrorIs(t, err, sutclient.ErrBadUser)
}

package user_test

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"lk_sut/internal/testutils"
	"net/http"
	"testing"
)

const (
	newPassword = "new_password"
)

func TestHandler_AddUser(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RegisterUserInitHandler()
	mocks.RegisterUserAuthHandler(testutils.Login, testutils.Password)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		RedisNil()

	mocks.RedisMock.
		ExpectHSet(mocks.Config.Redis.UserDataHTable, testutils.Login, testutils.Password).
		SetVal(1)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_AddUser_validation_error(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBodyWithLogin("bad_email")).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestHandler_AddUser_bad_user(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RegisterUserInitHandler()
	mocks.RegisterUserAuthHandlerBadUser(testutils.Login, testutils.Password)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		RedisNil()

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.BadUserResponse, resp.String())
}

func TestHandler_AddUser_user_exists(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal(testutils.Password)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.UserExistsResponse, resp.String())
}

func TestHandler_UpdateUser(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RegisterUserInitHandler()
	mocks.RegisterUserAuthHandler(testutils.Login, newPassword)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal(testutils.Password)

	mocks.RedisMock.
		ExpectHSet(mocks.Config.Redis.UserDataHTable, testutils.Login, newPassword).
		SetVal(1)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_UpdateUser_validation_error(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody("123")).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestHandler_UpdateUser_same_password(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(testutils.Password)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestHandler_UpdateUser_not_found(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		RedisNil()

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

func TestHandler_UpdateUser_wrong_old_password(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal("fake")

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

func TestHandler_UpdateUser_wrong_new_password(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RegisterUserInitHandler()
	mocks.RegisterUserAuthHandlerBadUser(testutils.Login, newPassword)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal(testutils.Password)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.BadUserResponse, resp.String())
}

func TestHandler_DeleteUser(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal(testutils.Password)

	mocks.RedisMock.
		ExpectHDel(mocks.Config.Redis.UserDataHTable, testutils.Login).
		SetVal(1)

	mocks.RedisMock.
		ExpectHDel(mocks.Config.Redis.UserLastLoginHTable, testutils.Login).
		SetVal(1)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Delete(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_DeleteUser_not_found(t *testing.T) {
	mocks := testutils.NewMockApp(t)

	t.Cleanup(mocks.ClearFunc)

	mocks.RedisMock.
		ExpectHGet(mocks.Config.Redis.UserDataHTable, testutils.Login).
		RedisNil()

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Delete(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

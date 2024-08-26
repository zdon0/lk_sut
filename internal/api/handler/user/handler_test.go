package user_test

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"lk_sut/internal/testutils"
)

const (
	newPassword = "new_password"
)

func TestHandler_AddUser(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetUser(testutils.Login, ""),

		testutils.RegisterUserInitHandler(),
		testutils.RegisterUserAuthHandler(testutils.Login, testutils.Password),

		testutils.ExpectCreateUser(testutils.Login, testutils.Password),
	}

	mocks := testutils.NewMockApi(t, expectations...)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_AddUser_validation_error(t *testing.T) {
	mocks := testutils.NewMockApi(t)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBodyWithLogin("bad_email")).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestHandler_AddUser_bad_user(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetUser(testutils.Login, ""),

		testutils.RegisterUserInitHandler(),
		testutils.RegisterUserAuthHandlerBadUser(testutils.Login, testutils.Password),
	}

	mocks := testutils.NewMockApi(t, expectations...)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.BadUserResponse, resp.String())
}

func TestHandler_AddUser_user_exists(t *testing.T) {
	mocks := testutils.NewMockApi(t, testutils.ExpectGetUser(testutils.Login, testutils.Password))

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Post(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.UserExistsResponse, resp.String())
}

func TestHandler_UpdateUser(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetUser(testutils.Login, testutils.Password),

		testutils.RegisterUserInitHandler(),
		testutils.RegisterUserAuthHandler(testutils.Login, newPassword),

		testutils.ExpectCreateUser(testutils.Login, newPassword),
	}

	mocks := testutils.NewMockApi(t, expectations...)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_UpdateUser_validation_error(t *testing.T) {
	mocks := testutils.NewMockApi(t)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody("123")).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
}

func TestHandler_UpdateUser_same_password(t *testing.T) {
	mocks := testutils.NewMockApi(t, testutils.ExpectGetUser(testutils.Login, testutils.Password))

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(testutils.Password)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_UpdateUser_not_found(t *testing.T) {
	mocks := testutils.NewMockApi(t, testutils.ExpectGetUser(testutils.Login, ""))

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

func TestHandler_UpdateUser_wrong_old_password(t *testing.T) {
	mocks := testutils.NewMockApi(t, testutils.ExpectGetUser(testutils.Login, "fake"))

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

func TestHandler_UpdateUser_wrong_new_password(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetUser(testutils.Login, testutils.Password),

		testutils.RegisterUserInitHandler(),
		testutils.RegisterUserAuthHandlerBadUser(testutils.Login, newPassword),
	}

	mocks := testutils.NewMockApi(t, expectations...)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserUpdateBody(newPassword)).
		Patch(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode())
	assert.JSONEq(t, testutils.BadUserResponse, resp.String())
}

func TestHandler_DeleteUser(t *testing.T) {
	expectations := []testutils.Opt{
		testutils.ExpectGetUser(testutils.Login, testutils.Password),
		testutils.ExpectDeleteUser(testutils.Login),
	}

	mocks := testutils.NewMockApi(t, expectations...)

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Delete(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode())
	assert.JSONEq(t, testutils.SimpleOkResponse, resp.String())
}

func TestHandler_DeleteUser_not_found(t *testing.T) {
	mocks := testutils.NewMockApi(t, testutils.ExpectGetUser(testutils.Login, ""))

	t.Cleanup(mocks.ClearFunc)

	resp, err := resty.New().R().
		SetBody(testutils.MakeUserBody()).
		Delete(mocks.Api.URL + "/api/v1/user")

	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode())
	assert.JSONEq(t, testutils.NotFoundResponse, resp.String())
}

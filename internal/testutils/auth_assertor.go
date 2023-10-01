package testutils

import (
	"fmt"
	"net/http"
	"testing"
)

type authAssertor struct {
	login    string
	password string
}

func newAuthAssertor(login string, password string) *authAssertor {
	return &authAssertor{
		login:    login,
		password: password,
	}
}

func (a *authAssertor) Assert(r *http.Request) error {
	reqLogin := r.FormValue("users")
	if reqLogin != a.login {
		return fmt.Errorf("bad login: expcted %s, got %s", a.login, reqLogin)
	}

	reqPassword := r.FormValue("parole")
	if reqPassword != a.password {
		return fmt.Errorf("bad password: expcted %s, got %s", a.password, reqPassword)
	}

	return nil
}

func (a *authAssertor) Log(t testing.TB) {
	t.Log("Testing request for required user")
}

func (a *authAssertor) Error(t testing.TB, err error) {
	t.Errorf("assertion error: %s", err)
}

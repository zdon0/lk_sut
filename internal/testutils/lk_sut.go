package testutils

import "net/http"

func RegisterUserInitHandler() Opt {
	return func(app *App) {
		app.sutServer.NewHandler().Get("/cabinet").
			AssertHeaderValue("Cookie", "").
			Reply(http.StatusOK).
			AddHeader(SetCookieHeader, UidCookieHeaderResponse())
	}
}

func RegisterUserAuthHandler(login string, password string) Opt {
	return func(app *App) {
		app.sutServer.NewHandler().
			Post("/cabinet/lib/autentificationok.php").
			AssertCustom(newAuthAssertor(login, password)).
			AssertHeaderValue(CookieHeader, UidCookieHeaderRequest()).
			Reply(http.StatusOK).
			BodyString("1")
	}
}

func RegisterUserAuthHandlerBadUser(login, password string) Opt {
	return func(app *App) {
		app.sutServer.NewHandler().
			Post("/cabinet/lib/autentificationok.php").
			AssertCustom(newAuthAssertor(login, password)).
			AssertHeaderValue(CookieHeader, UidCookieHeaderRequest()).
			Reply(http.StatusOK).
			BodyString("0")
	}
}

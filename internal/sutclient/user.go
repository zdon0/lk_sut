package sutclient

import (
	"context"
	"net/url"
)

func (sc *SutClient) initUser(ctx context.Context) error {
	sc.resetCookie()

	// Set cookie session
	_, err := sc.r(ctx).
		Get("/cabinet")

	return err
}

func (sc *SutClient) authUser(ctx context.Context, login, password string) error {
	formData := make(url.Values, 2)

	formData.Set("users", login)
	formData.Set("parole", password)

	_, err := sc.r(ctx).
		SetFormDataFromValues(formData).
		Post("/cabinet/lib/autentificationok.php")

	return err
}

func (sc *SutClient) AuthorizeUser(ctx context.Context, login, password string) error {
	// goroutine user validation in handlers
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	if err := sc.initUser(ctx); err != nil {
		return err
	}

	return sc.authUser(ctx, login, password)
}

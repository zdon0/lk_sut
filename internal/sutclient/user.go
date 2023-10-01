package sutclient

import (
	"context"
	"lk_sut/internal/domain"
	"net/url"
)

func (sc *SutClient) initUser(ctx context.Context) error {
	sc.resetCookie()

	// Set cookie session
	_, err := sc.r(ctx).
		Get("/cabinet")

	return err
}

func (sc *SutClient) authUser(ctx context.Context, user domain.User) error {
	formData := make(url.Values, 2)

	formData.Set("users", user.Login)
	formData.Set("parole", user.Password)

	_, err := sc.r(ctx).
		SetFormDataFromValues(formData).
		Post("/cabinet/lib/autentificationok.php")

	return err
}

func (sc *SutClient) AuthorizeUser(ctx context.Context, user domain.User) error {
	// goroutine user validation in handlers
	sc.mutex.Lock()
	defer sc.mutex.Unlock()

	if err := sc.initUser(ctx); err != nil {
		return err
	}

	return sc.authUser(ctx, user)
}

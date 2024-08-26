package sutclient

import (
	"net/http"

	"github.com/go-resty/resty/v2"

	"lk_sut/internal/pkg/decoder"
)

const (
	forbiddenResponse = "У Вас нет прав доступа. Или необходимо перезагрузить приложение.."
)

func afterResponse(_ *resty.Client, r *resty.Response) error {
	if !r.IsSuccess() {
		return ErrUnexpectedStatusCode
	}

	url := r.Request.RawRequest.URL

	switch url.Path {
	case "/cabinet":
		if r.StatusCode() != http.StatusOK {
			return ErrFailedToInitUser
		}

		return nil

	case "/cabinet/lib/autentificationok.php":
		if r.String() != "1" {
			return ErrBadUser
		}

		return nil

	case "/cabinet/project/cabinet/forms/raspisanie.php":
		resp, err := decoder.Decode(r.String())
		if err != nil {
			return err
		}

		if resp == forbiddenResponse {
			return ErrForbidden
		}

		return nil
	}

	return nil
}

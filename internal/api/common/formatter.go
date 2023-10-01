package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"lk_sut/internal/domain"
	"lk_sut/pkg/dto"
	"net/http"
)

func MakeSimpleOkResponse(c *gin.Context) {
	resp := dto.SimpleOkResponse{
		Result: dto.SimpleOkResult{
			Status: "ok",
		},
	}

	c.JSON(http.StatusOK, resp)
}

func MakeErrorResponse(c *gin.Context, err error) {
	_ = c.Error(err)

	resp := dto.ErrorResponse{
		Error: err.Error(),
	}

	switch err.(type) {
	case validation.Errors:
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	switch {
	case errors.Is(err, domain.ErrBadUser), errors.Is(err, domain.ErrUserExists):
		c.JSON(http.StatusBadRequest, resp)
		return
	case errors.Is(err, domain.ErrNotFound):
		c.JSON(http.StatusNotFound, resp)
		return
	}

	c.JSON(http.StatusInternalServerError, resp)
}

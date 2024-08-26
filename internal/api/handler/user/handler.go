package user

import (
	"github.com/gin-gonic/gin"

	"lk_sut/internal/api/common"
	"lk_sut/internal/interactor/user"
	"lk_sut/pkg/dto"
)

type Handler struct {
	interactor *user.Interactor
}

func NewHandler(userInteractor *user.Interactor) *Handler {
	return &Handler{
		interactor: userInteractor,
	}
}

// AddUser 		 godoc
// @Summary      Add user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request body dto.User true "user info"
// @Success      200  {object}  dto.SimpleOkResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/user [post]
func (h *Handler) AddUser(c *gin.Context) {
	var req dto.User

	if err := c.BindJSON(&req); err != nil {
		return
	}

	if err := h.interactor.AddUser(c, makeUserDomain(req)); err != nil {
		common.MakeErrorResponse(c, err)
		return
	}

	common.MakeSimpleOkResponse(c)
}

// UpdateUser	 godoc
// @Summary      Update password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request body dto.UpdateUser true "user info"
// @Success      200  {object}  dto.SimpleOkResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/user [patch]
func (h *Handler) UpdateUser(c *gin.Context) {
	var req dto.UpdateUser

	if err := c.BindJSON(&req); err != nil {
		return
	}

	if err := h.interactor.UpdateUser(c, makeUpdateUserDomain(req)); err != nil {
		common.MakeErrorResponse(c, err)
		return
	}

	common.MakeSimpleOkResponse(c)
}

// DeleteUser	 godoc
// @Summary      Delete user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request body dto.User true "user info"
// @Success      200  {object}  dto.SimpleOkResponse
// @Failure      400  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Failure      500  {object}  dto.ErrorResponse
// @Router       /api/v1/user [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	var req dto.User

	if err := c.BindJSON(&req); err != nil {
		return
	}

	if err := h.interactor.DeleteUser(c, makeUserDomain(req)); err != nil {
		common.MakeErrorResponse(c, err)
		return
	}

	common.MakeSimpleOkResponse(c)
}

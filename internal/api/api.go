package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "lk_sut/docs" // gin swagger
	"lk_sut/internal/api/handler/user"
	"lk_sut/internal/config"
)

func NewApi(cfg *config.Config, logger *zap.Logger, userHandler *user.Handler) *gin.Engine {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := newGinEngine(logger)

	apiV1group := handler.Group("/api/v1")

	registerDefaultRoutes(handler)
	registerUserRoutes(apiV1group, userHandler)

	return handler
}

func newGinEngine(logger *zap.Logger) *gin.Engine {
	r := gin.New()

	r.Use(defaultZapLogger(logger, "/swagger", "/ready"))
	r.Use(defaultZapRecovery(logger, true))

	r.ContextWithFallback = true

	return r
}

func registerDefaultRoutes(r *gin.Engine) {
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.DefaultModelsExpandDepth(-1),
		),
	)

	r.GET("/ready", simpleOkHandler)
}

func registerUserRoutes(api *gin.RouterGroup, h *user.Handler) {
	api.POST("/user", h.AddUser)
	api.PATCH("/user", h.UpdateUser)
	api.DELETE("/user", h.DeleteUser)
}

func simpleOkHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

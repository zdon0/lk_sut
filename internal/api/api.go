package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	_ "lk_sut/docs" // gin swagger
	"lk_sut/internal/config"
	"lk_sut/internal/env"
	"net/http"
)

func NewApi(cfg *config.Config, e *env.Env, logger *zap.Logger) *http.Server {
	addr := fmt.Sprintf("%s:%d", cfg.Api.Addr, cfg.Api.Port)

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := newGinEngine(logger)

	apiV1group := handler.Group("/api/v1")

	registerDefaultRoutes(handler)
	registerUserRoutes(apiV1group, e)

	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: cfg.Api.ReadHeaderTimeout,
	}
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

func registerUserRoutes(api *gin.RouterGroup, e *env.Env) {
	api.POST("/user", e.User.AddUser)
	api.PATCH("/user", e.User.UpdateUser)
	api.DELETE("/user", e.User.DeleteUser)
}

func simpleOkHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

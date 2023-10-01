package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func defaultZapLogger(logger *zap.Logger, pathPrefixes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		for _, pathPrefix := range pathPrefixes {
			if strings.HasPrefix(path, pathPrefix) {
				return
			}
		}

		start := time.Now().UTC()

		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.Duration("latency", latency),
		}

		if len(c.Errors) > 0 || c.IsAborted() || c.Writer.Status() >= 400 {
			if len(c.Errors) > 0 {
				fields = append(fields, zap.Strings("error", c.Errors.Errors()))
			}

			logger.Error("Gin Error", fields...)
		} else {
			logger.Info("Gin Success", fields...)
		}
	}
}

func defaultZapRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						brokenPipe = isBrokenPipe(se)
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, true)

				fields := []zap.Field{
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.ByteString("request", httpRequest),
				}

				if brokenPipe {
					logger.Error("Gin Recovery from panic", fields...)

					_ = c.Error(err.(error))
					c.Abort()

					return
				}

				if errError, ok := err.(error); ok {
					_ = c.Error(errError)
				}

				if stack {
					fields = append(fields, zap.ByteString("stack", debug.Stack()))
				}

				logger.Error("Gin Recovery from panic", fields...)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

func isBrokenPipe(err *os.SyscallError) bool {
	errString := strings.ToLower(err.Error())

	return strings.Contains(errString, "broken pipe") || strings.Contains(errString, "connection reset by peer")
}

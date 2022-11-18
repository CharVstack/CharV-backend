package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/skerkour/rz"
)

func Logger(logger rz.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		if err := c.Errors.ByType(gin.ErrorTypePrivate).Last(); err != nil {
			logger.Error(
				err.Error(),
				rz.Int("status", c.Writer.Status()),
				rz.String("method", c.Request.Method),
				rz.String("path", c.Request.URL.Path),
				rz.String("query", c.Request.URL.RawQuery),
				rz.String("ip", c.ClientIP()),
				rz.String("user-agent", c.Request.UserAgent()),
				rz.Duration("elapsed", time.Since(start)),
			)
		} else if err := c.Errors.ByType(gin.ErrorTypePublic).Last(); err != nil {
			logger.Warn(
				err.Error(),
				rz.Int("status", c.Writer.Status()),
				rz.String("method", c.Request.Method),
				rz.String("path", c.Request.URL.Path),
				rz.String("query", c.Request.URL.RawQuery),
				rz.String("ip", c.ClientIP()),
				rz.String("user-agent", c.Request.UserAgent()),
				rz.Duration("elapsed", time.Since(start)),
			)
		} else {
			logger.Info(
				"Logger",
				rz.Int("status", c.Writer.Status()),
				rz.String("method", c.Request.Method),
				rz.String("path", c.Request.URL.Path),
				rz.String("query", c.Request.URL.RawQuery),
				rz.String("ip", c.ClientIP()),
				rz.String("user-agent", c.Request.UserAgent()),
				rz.Duration("elapsed", time.Since(start)),
			)
		}
	}
}

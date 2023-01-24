package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GenericErrorHandler(c *gin.Context, err error, statusCode int) {
	if statusCode >= 400 && statusCode < 500 {
		err = c.Error(err).SetType(gin.ErrorTypePublic)
		c.AbortWithStatusJSON(statusCode, gin.H{"message": err.Error()})
	} else if statusCode >= 500 && statusCode < 600 {
		err = c.Error(err).SetType(gin.ErrorTypePrivate)
		c.AbortWithStatus(statusCode)
	}
}

func ValidationErrorHandler(c *gin.Context, message string, statusCode int) {
	err := errors.New(message)
	GenericErrorHandler(c, err, statusCode)
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "login") {
			c.Next()
			return
		}
	}
}

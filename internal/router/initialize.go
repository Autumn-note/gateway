package router

import (
	"gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Interceptor())

	userGroup := Router.Group("")
	NewUserRouter().InitUserRouter(userGroup)

	return Router
}

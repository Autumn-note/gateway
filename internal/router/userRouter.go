package router

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (u *UserRouter) UserRouter(router *gin.RouterGroup) gin.IRoutes {
	user := router.Group("user")

	{
	}

	return user
}

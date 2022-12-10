package router

import (
	"gateway/internal/grpc/client"
	"gateway/internal/handler"
	"gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) gin.IRoutes {
	user := router.Group("/user").Use(middleware.User())
	userClient, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	userHandler := handler.UserHandler{
		Client: userClient,
	}
	{
		user.GET("/get", userHandler.GetUser)
	}

	return user
}

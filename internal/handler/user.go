package handler

import (
	"gateway/internal/grpc/client"
	user "gateway/internal/grpc/pb"
	"github.com/gin-gonic/gin"
	"log"
)

type UserHandler struct {
	Client *client.UserClient
}

func (u *UserHandler) InitHandler() {

}

func (u *UserHandler) GetUser(c *gin.Context) {
	req := &user.GetUserRequest{}
	userResponse, err := u.Client.GetUser(c, req)
	if err != nil {
		log.Println("fail to get resp from client %w", err)
		c.JSON(500, err)
	}
	log.Print("userResponse: ", userResponse)
	c.JSON(200, toUserResponse(userResponse))
}

type UserResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func toUserResponse(user *user.GetUserResponse) *UserResponse {
	return &UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}
}

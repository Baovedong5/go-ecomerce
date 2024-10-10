package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuong/go-ecomerce/internal/service"
	"github.com/phuong/go-ecomerce/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller -> service -> repo -> model -> db
func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessReponse(c, 20001, []string{"phuong", "thanh"})
}

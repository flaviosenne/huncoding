package controller

import (
	"github.com/flaviosenne/huncoding/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByEmail(*gin.Context)
	FindUserById(*gin.Context)

	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
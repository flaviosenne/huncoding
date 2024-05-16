package controller

import (
	"github.com/flaviosenne/huncoding/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainServiceInterface) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	FindUserByEmail(*gin.Context)
	FindUserByID(*gin.Context)

	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainServiceInterface
}

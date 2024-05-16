package service

import (
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository"
)

func NewUserDomainService(repo repository.UserRepositoryInterface) UserDomainServiceInterface {
	return &userDomainService{
		repository: repo,
	}
}

type userDomainService struct {
	repository repository.UserRepositoryInterface
}

type UserDomainServiceInterface interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

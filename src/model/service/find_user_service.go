package service

import (
	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicio da busca do usuário por id service",
		zap.String("journey", "findUserByID"))

	return ud.repository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicio da busca do usuário por email service",
		zap.String("journey", "findUserByEmail"))
	return ud.repository.FindUserByEmail(email)
}

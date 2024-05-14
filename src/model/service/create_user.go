package service

import (
	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicio da criação do usuário service",
		zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	userDomainRepository, err := ud.repository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Erro em criar usuário service", err,
			zap.String("journey", "createUser"))
		return nil, err
	}
	logger.Info("Criado usuário com sucesso service",
		zap.String("journey", "createUser"),
		zap.String("userId", userDomainRepository.GetID()))
	return userDomainRepository, nil
}

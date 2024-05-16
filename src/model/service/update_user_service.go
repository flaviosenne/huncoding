package service

import (
	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Inicio da criação do usuário service",
		zap.String("journey", "createUser"))

	if err := ud.repository.UpdateUser(id, userDomain); err != nil {
		logger.Error("Erro em atualizar usuário service", err,
			zap.String("journey", "createUser"))
		return err
	}
	logger.Info("Atualizado usuário com sucesso service",
		zap.String("journey", "updateUser"),
		zap.String("userId", id))
	return nil
}

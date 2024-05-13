package model

import (
	"fmt"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Inicio da criação do usuário domínio", zap.String("journey", "createUser"))
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}

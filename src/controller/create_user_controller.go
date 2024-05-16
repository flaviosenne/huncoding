package controller

import (
	"net/http"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/validation"
	"github.com/flaviosenne/huncoding/src/controller/model/request"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Iniciando criação de usuário controller",
		zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro em tentar converter objeto", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	result, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error("Erro em criar usuário  controller", err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Criado usuário com sucesso controller",
		zap.String("userId", domain.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(result))
}

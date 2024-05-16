package controller

import (
	"net/http"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/validation"
	"github.com/flaviosenne/huncoding/src/controller/model/request"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {

	logger.Info("Iniciando atualização de usuário controller",
		zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro em tentar converter objeto", err,
			zap.String("journey", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	userId := c.Param("id")

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Erro em atualizar usuário controller", err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Atualizado usuário com sucesso controller",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusNoContent)
}

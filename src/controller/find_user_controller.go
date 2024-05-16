package controller

import (
	"net/http"
	"net/mail"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {

	logger.Info("Incio da busca do usuário por id controller",
		zap.String("journey", "findUserByID"))

	userId := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError("userId não é um UUID válido")
		logger.Error("id do usuário não é um UUID válido controller", err,
			zap.String("userId", userId),
			zap.String("journey", "findUserByID"))
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		logger.Error("Erro em buscar usuário pelo id controller", err,
			zap.String("userId", userId),
			zap.String("journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Sucesso em buscar usuário pelo id controller",
		zap.String("userId", userId),
		zap.String("journey", "findUserByID"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	logger.Info("Incio da busca do usuário por email controller",
		zap.String("journey", "findUserByEmail"))

	userEmail := c.Param("email")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("email não é válido")
		logger.Error("email do usuário não é válido controller", err,
			zap.String("email", userEmail),
			zap.String("journey", "findUserByEmail"))
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		logger.Error("Erro em buscar usuário pelo email controller", err,
			zap.String("email", userEmail),
			zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Sucesso em buscar usuário pelo email controller",
		zap.String("email", userEmail),
		zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

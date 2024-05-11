package controller

import (
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro em criar usuário")
	c.JSON(err.Code, err)
}

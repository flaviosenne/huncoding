package controller

import (
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {

}

func FindUserByEmail(c *gin.Context) {
	err := rest_err.NewBadRequestError("Email não existe")
	c.JSON(err.Code, err)
}

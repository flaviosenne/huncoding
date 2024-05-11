package controller

import (
	"fmt"

	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Os campos estão incorretos, error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

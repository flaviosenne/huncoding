package controller

import (
	"fmt"

	"github.com/flaviosenne/huncoding/src/configuration/validation"
	"github.com/flaviosenne/huncoding/src/model/request"
	"github.com/flaviosenne/huncoding/src/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(200, response)
}

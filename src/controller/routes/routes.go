package routes

import (
	"github.com/flaviosenne/huncoding/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", controller.FindUserById)
	r.GET("/user-email/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}

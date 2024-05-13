package main

import (
	"log"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/controller"
	"github.com/flaviosenne/huncoding/src/controller/routes"
	"github.com/flaviosenne/huncoding/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Está começando")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro em carregar variáveis de ambiente")
	}

	//inicializar as dependencias
	sevice := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(sevice)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

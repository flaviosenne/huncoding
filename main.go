package main

import (
	"context"
	"log"

	"github.com/flaviosenne/huncoding/src/configuration/database/mongodb"
	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/controller"
	"github.com/flaviosenne/huncoding/src/controller/routes"
	"github.com/flaviosenne/huncoding/src/model/repository"
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
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatal("Erro em conectar com banco de dados")
	}
	repo := repository.NewUserRepository(database)
	sevice := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(sevice)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

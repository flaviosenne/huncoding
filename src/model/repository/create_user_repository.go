package repository

import (
	"context"
	"os"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Iniciou criação do usuário no repository", zap.String("journey", "createUser"))
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConverterDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Erro em salvar usuário nabase de dados repository", err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("Iniciou criação do usuário no repository",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConverterEntityToDomain(*value), nil
}

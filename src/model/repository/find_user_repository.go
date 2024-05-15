package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository/entity"
	"github.com/flaviosenne/huncoding/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciou busca do usuário por email no repository",
		zap.String("journey", "findUserByEmail"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Usuário não encontrado com email: %s", email)

			logger.Error(errorMessage, err,
				zap.String("journey", "findUserByEmail"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro em tentar buscar usuário por email"

		logger.Error(errorMessage, err,
			zap.String("journey", "findUserByEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("Sucesso em buscar usuário por email no repository",
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserByEmail"))

	return converter.ConverterEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciou busca do usuário por ID no repository",
		zap.String("journey", "findUserByID"))

	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Usuário não encontrado com ID: %s", id)

			logger.Error(errorMessage, err,
				zap.String("journey", "findUserById"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro em tentar buscar usuário por ID"

		logger.Error(errorMessage, err,
			zap.String("journey", "findUserById"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("Sucesso em buscar usuário por ID no repository",
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("journey", "findUserById"))

	return converter.ConverterEntityToDomain(*userEntity), nil

}

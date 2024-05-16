package repository

import (
	"context"
	"os"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"github.com/flaviosenne/huncoding/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Iniciou criação do usuário no repository",
		zap.String("journey", "updateUser"))
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConverterDomainToEntity(userDomain)

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Erro em tentar converter id para tipo HEX", err,
			zap.String("journey", "updateUser"))
		return rest_err.NewBadRequestError("id do usuário não é do tipo válido")
	}
	filter := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: value}}

	if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
		logger.Error("Erro em atualizar usuário na abase de dados repository", err,
			zap.String("journey", "updateUser"))

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("Atualização com sucesso do usuário no repository",
		zap.String("userId", id),
		zap.String("journey", "updateUser"))

	return nil
}

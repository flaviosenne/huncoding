package repository

import (
	"context"
	"os"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Iniciou criação do usuário no repositório")
	collectionName := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collectionName)
	value, err := userDomain.GetJsonValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))
	return userDomain, nil
}

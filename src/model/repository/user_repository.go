package repository

import (
	"github.com/flaviosenne/huncoding/src/configuration/rest_err"
	"github.com/flaviosenne/huncoding/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepositoryInterface {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
}

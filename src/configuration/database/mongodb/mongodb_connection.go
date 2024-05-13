package mongodb

import (
	"context"
	"os"

	"github.com/flaviosenne/huncoding/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	mongodDatabase := os.Getenv(MONGODB_USER_DB)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Conseguiu conectar mongodb")
	return client.Database(mongodDatabase), nil
}

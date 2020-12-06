package db

import (
	"context"
	"github.com/ginkt/mtuci_backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(ctx context.Context, cfg *config.Config) (client *mongo.Client, err error) {
	credential := options.Credential{
		Username: cfg.MongoUsername,
		Password: cfg.MongoPassword,
	}
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/").SetAuth(credential)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return
	}
	return
}

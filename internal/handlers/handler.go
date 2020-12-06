package handlers

import "go.mongodb.org/mongo-driver/mongo"

type Handler struct {
	mongo *mongo.Client
}

func CreateHandler(mongoClient *mongo.Client) *Handler {
	return &Handler{mongo: mongoClient}
}

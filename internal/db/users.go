package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRecord struct {
	Id                primitive.ObjectID `bson:"_id"`
	Name              string             `bson:"name" ,json:"name"`
	Age               int64             `bson:"age" ,json:"age"`
	DrivingExperience int64              `bson:"drivingExperience" ,json:"drivingExperience"`
	Sex               string             `bson:"sex" ,json:"sex"`
	MobileDevice      string             `bson:"mobileDevice" ,json:"mobileDevice"`
}

func InsertNewUserRecordInUsers(mongoConn *mongo.Client, ctx context.Context, user UserRecord) (insertedId primitive.ObjectID, err error) {
	collection := mongoConn.Database("mtuciBackend").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedId = result.InsertedID.(primitive.ObjectID)
	return
}

func IsUserExists(mongoConn *mongo.Client, ctx context.Context, userId primitive.ObjectID) (exists bool, err error) {
	collection := mongoConn.Database("mtuciBackend").Collection("users")
	result := collection.FindOne(ctx, bson.M{ "_id": userId})
	if result.Err() == mongo.ErrNoDocuments {
		return false, nil
	}
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}

package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RideRecord struct {
	Id                primitive.ObjectID `bson:"_id"`
	UserId 			  primitive.ObjectID `bson:"userId" ,json:"userId"`
	Position string `bson:"position" ,json:"position"`
	RoadState string `bson:"roadState" ,json:"roadState"`
	TimeOfADay	string `bson:"timeOfADay" ,json:"timeOfADay"`
	Wheels string `bson:"wheels" ,json:"wheels"`
}

func InsertNewRideRecordInRides(mongoConn *mongo.Client, ctx context.Context, ride RideRecord) (insertedId primitive.ObjectID, err error) {
	collection := mongoConn.Database("mtuciBackend").Collection("rides")
	result, err := collection.InsertOne(ctx, ride)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedId = result.InsertedID.(primitive.ObjectID)
	return
}

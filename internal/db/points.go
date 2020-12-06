package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type PointRecord struct {
	Id             primitive.ObjectID `bson:"_id"`
	RideId         primitive.ObjectID `bson:"rideId" ,json:"rideId"`
	Longitude      float64            `bson:"longitude" ,json:"longitude"` // долгота
	Latitude       float64            `bson:"latitude" ,json:"latitude"`   // широта
	Speed          float64            `bson:"speed" ,json:"speed"`
	AccelerometerX float64            `bson:"accelerometerX" ,json:"accelerometerX"`
	AccelerometerY float64            `bson:"accelerometerY" ,json:"accelerometerY"`
	AccelerometerZ float64            `bson:"accelerometerZ" ,json:"accelerometerZ"`
	GyroscopeX     float64            `bson:"gyroscopeX" ,json:"gyroscopeX"`
	GyroscopeY     float64            `bson:"gyroscopeY" ,json:"gyroscopeY"`
	GyroscopeZ     float64            `bson:"gyroscopeZ" ,json:"gyroscopeZ"`
	Timestamp      time.Time          `bson:"timestamp" ,json:"timestamp"`
}

func AddNewPointRecordToPoints(mongoConn *mongo.Client, ctx context.Context, point PointRecord) (insertedId primitive.ObjectID, err error) {
	collection := mongoConn.Database("mtuciBackend").Collection("points")
	result, err := collection.InsertOne(ctx, point)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	insertedId = result.InsertedID.(primitive.ObjectID)
	return
}
package durable

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func SaveIdentity(id int64) bool {
	uri := "mongodb://dragon:dragon@localhost:27017/dragon"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("dragon").Collection("clients")
	doc := bson.D{{"id", id}, {"name", "Ayden Franklin"}}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
	return result != nil
}
package configsdb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// Thiết lập thông tin kết nối MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Lấy collection trong MongoDB

	// collection_admin = client.Database("DB-test").Collection("admin")
	return client

}

var client *mongo.Client = ConnectDB()

func GetCollection(collectioneName string) *mongo.Collection {
	collection := client.Database("DB-test").Collection("collectioneName")
	return collection
}

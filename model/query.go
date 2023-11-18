package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func QueryCreate(collection *mongo.Collection, Newuser User) (User, error) {
	_, err := collection.InsertOne(context.Background(), Newuser)
	return Newuser, err

}
func QueryGetbyID(collection *mongo.Collection, id string) (User, error) {
	var user User
	filter := bson.M{"userid": id}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	return user, err
}
func QueryEdit(collection *mongo.Collection, id string, updateData bson.M) (*mongo.UpdateResult, error) {
	filter := bson.M{"userid": id}
	update := bson.M{"$set": updateData}
	var query, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return query, nil
}
func QueryDetele(collection *mongo.Collection, id string) (*mongo.DeleteResult, error) {
	query, err := collection.DeleteOne(context.TODO(), bson.M{"userid": id})
	if err != nil {
		return nil, err
	}
	return query, nil

}

package model

import (
	"App-Task/configsdb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configsdb.GetCollection("users")

func QueryeEditUser(c *gin.Context, Newuser User, userId string) {
	update := bson.M{"fullname": Newuser.Fullname, "email": Newuser.Email, "phone": Newuser.Phone, "address": Newuser.Address, "role": Newuser.Role}
	query, err := userCollection.UpdateOne(context.TODO(), bson.M{"userid": userId}, bson.M{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if query.ModifiedCount > 0 {
		c.JSON(http.StatusOK, gin.H{"Message": "Update access"})
	}
}

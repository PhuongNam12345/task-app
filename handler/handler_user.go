package Handler

import (
	"App-Task/configsdb"
	"App-Task/model"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configsdb.GetCollection("users")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var Newuser model.User

		if err := c.BindJSON(&Newuser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error()})
			return
		}
		id := Newuser.UserID
		fmt.Println(id)
		var result bson.M
		err := userCollection.FindOne(context.TODO(), bson.M{"userid": id}).Decode(&result)
		if err != nil {
			_, err := userCollection.InsertOne(context.TODO(), Newuser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Messeger": "Create user access!"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Messeger": "Id exists!"})

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		var Newuser model.User
		err := userCollection.FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&Newuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Newuser)
	}
}
func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		var Newuser model.User
		if err := c.BindJSON(&Newuser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
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
}
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		_, err := userCollection.DeleteMany(context.TODO(), bson.M{"userid": userId})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, "Delete access")
	}
}

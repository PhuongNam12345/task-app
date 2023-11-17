package Handler

import (
	"App-Task/configsdb"
	"App-Task/model"
	"context"
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
		_, err := userCollection.InsertOne(context.TODO(), Newuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Messeger": "Create user access!"})

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		var Newuser model.User
		err := userCollection.FindOne(context.TODO(), bson.M{"userid": userId}).Decode(&Newuser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, &Newuser)

	}
}

package Handler

import (
	"App-Task/configsdb"
	"App-Task/model"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userCollection = configsdb.GetCollection("users")
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

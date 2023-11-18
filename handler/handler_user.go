package Handler

import (
	"App-Task/configsdb"
	"App-Task/model"
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
		_, err := model.QueryGetbyID(userCollection, id)
		if err != nil {
			_, err := model.QueryCreate(userCollection, Newuser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Messeger": "Create user access!", "User": Newuser})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Messeger": "Id exists!"})

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		resultQuery, err := model.QueryGetbyID(userCollection, userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &resultQuery)
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

		resultQuery, err := model.QueryEdit(userCollection, userId, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if resultQuery.ModifiedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"Message": "Update access"})
		}
	}
}
func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userid")
		resultQuery, err := model.QueryDetele(userCollection, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if resultQuery.DeletedCount > 0 {
			c.JSON(http.StatusOK, "Deleted successfully!")
		}
	}
}

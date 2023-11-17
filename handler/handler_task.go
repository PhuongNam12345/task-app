/* trunk-ignore-all(gofmt) */
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

var taskCollection *mongo.Collection = configsdb.GetCollection("tasks")

func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {

		var Newtask model.Task

		if err := c.BindJSON(&Newtask); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"errorss": err.Error()})

			return
		}
		id := Newtask.TaskID
		var result bson.M
		err := taskCollection.FindOne(context.TODO(), bson.M{"taskid": id}).Decode(&result)
		if err != nil {
			_, err := taskCollection.InsertOne(context.TODO(), Newtask)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Messeger": "Create task access!", "Task": Newtask})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"Messeger": "Id task exists!"})

	}
}
func GetTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("taskid")
		var Newtask model.Task
		err := taskCollection.FindOne(context.TODO(), bson.M{"taskid": taskId}).Decode(&Newtask)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Newtask)
	}
}

func EditTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("taskid")
		var Newtask model.Task
		if err := c.BindJSON(&Newtask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		update := bson.M{"taskname": Newtask.Taskname, "starting": Newtask.Starting, "deadline": Newtask.Deadline, "catelogy": Newtask.Catelogy, "userid": Newtask.UserID}
		query, err := taskCollection.UpdateOne(context.TODO(), bson.M{"taskid": taskId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if query.ModifiedCount > 0 {
			c.JSON(http.StatusOK, gin.H{"Message": "Update task access", "Task": Newtask})
		}
	}
}
func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("taskid")
		_, err := taskCollection.DeleteOne(context.TODO(), bson.M{"taskid": taskId})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, "Delete task access")
	}
}

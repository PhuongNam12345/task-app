package Handler

import (
	"App-Task/configsdb"

	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection = configsdb.GetCollection("tasks")

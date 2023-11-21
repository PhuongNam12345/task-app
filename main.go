package main

import (
	"App-Task/configsdb"
	"App-Task/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//run database
	configsdb.ConnectDB()
	// configsdb.GetCollection("users")
	// configsdb.GetCollection("tasks")
	routes.UserRouter(router)
	routes.TaskRouter(router)
	router.Run("localhost:5000")
}

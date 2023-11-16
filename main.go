package main

import (
	"App-Task/configsdb"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//run database
	configsdb.ConnectDB()
	configsdb.GetCollection("users")
	router.Run("localhost:5000")
}

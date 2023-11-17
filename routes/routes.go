package routes

import (
	Handler "App-Task/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

	router.POST("/users", Handler.CreateUser())
}

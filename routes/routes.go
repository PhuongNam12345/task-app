package routes

import (
	Handler "App-Task/handler"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	router.POST("/users", Handler.CreateUser())
	router.GET("/users/:userid", Handler.GetUser())
	router.PUT("/users/:userid", Handler.EditUser())
	router.DELETE("/users/:userid", Handler.DeleteUser())
}

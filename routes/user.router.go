package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/madhavarora03/golang-jwt-project/controllers"
	"github.com/madhavarora03/golang-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

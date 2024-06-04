package routes

import (
	"github.com/gin-gonic/gin"
	"user-service/controllers"
)

func SetupRouter() *gin.Engine {
	// Set up the routes
	router := gin.Default()
	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/signup", controllers.SignUp)
	}
	return router
}

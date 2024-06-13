package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/controllers"
	"user-service/middleware"
	"user-service/services"
)

func SetupRouter() *gin.Engine {
	// Set up the routes
	router := gin.Default()
	jwtService := services.NewJWTService()
	api := router.Group("/api/users")
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "Users API v1"})
		})
		api.POST("/signup", controllers.SignUp)
		api.POST("/signin", controllers.SignIn)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(jwtService))
		{
			//protected.POST("/signup", controllers.SignUp)
		}
	}
	return router
}

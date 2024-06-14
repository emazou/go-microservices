package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/controllers"
	"user-service/middlewares"
	"user-service/services"
	"user-service/utils"
)

func SetupRouter() *gin.Engine {
	// Set up the routes
	router := gin.Default()
	jwtService := services.NewJWTService()
	api := router.Group(utils.BaseUrl)
	{
		api.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "Users API v1"})
		})
		api.POST(utils.SignupUrl, controllers.SignUp)
		api.POST(utils.SigninUrl, controllers.SignIn)
		api.DELETE("/:id", controllers.DeleteUserByID)
	}
	protected := api.Group("/")
	protected.Use(middlewares.AuthMiddleware(jwtService))
	{
		protected.DELETE("/:id", controllers.DeleteUserByID)
	}
	return router
}

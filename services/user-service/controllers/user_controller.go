package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/services"
)

type CreateUserInput struct {
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	Address       string `json:"address" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	Role          string `json:"role" binding:"required"`
}

// SignUp creates a new user account in the system with the provided details
func SignUp(c *gin.Context) {
	// Create a new user
	var user CreateUserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userCreated, err := services.SignUpService(
		user.Name,
		user.LastName,
		user.Email,
		user.Password,
		user.Address,
		user.PaymentMethod,
		user.Role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": userCreated})
}

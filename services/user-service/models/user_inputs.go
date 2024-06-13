package models

type SignUpUserInput struct {
	Name          string        `json:"name" binding:"required"`
	LastName      string        `json:"last_name" binding:"required"`
	Email         string        `json:"email" binding:"required,email"`
	Password      string        `json:"password" binding:"required"`
	Address       string        `json:"address" binding:"required"`
	PaymentMethod PaymentMethod `json:"payment_method" binding:"required"`
	Role          Role          `json:"role" binding:"required"`
}

type SignInUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

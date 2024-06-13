package models

type UserResponse struct {
	Name          string        `json:"name"`
	LastName      string        `json:"last_name"`
	Email         string        `json:"email"`
	Address       string        `json:"address"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	Role          Role          `json:"role"`
}

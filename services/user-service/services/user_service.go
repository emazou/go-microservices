package services

import (
	"user-service/models"
	"user-service/repositories"
)

// SignUpService creates a new user in the database and returns the user and the error if any occurred
func SignUpService(name, lastName, email, password, address, paymentMethod, role string) (*models.User, error) {
	// Create a new user
	user := &models.User{
		Name:          name,
		LastName:      lastName,
		Email:         email,
		Password:      password,
		Address:       address,
		PaymentMethod: models.PaymentMethod(paymentMethod),
		Role:          models.Role(role),
	}
	// Create the user
	err := repositories.CreateUser(user)
	if err != nil {
		return user, err
	}
	user, err = repositories.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}

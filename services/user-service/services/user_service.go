package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"user-service/models"
	"user-service/repositories"
)

// SignUpService creates a new user in the database and returns the user and the error if any occurred
func SignUpService(name, lastName, email, password, address, paymentMethod, role string) (*models.User, error) {
	// Define the email regex pattern to validate the email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched := regexp.MustCompile(emailRegex).MatchString(email); !matched {
		return nil, errors.New("invalid email address")
	}
	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// Create a new user
	user := &models.User{
		Name:          name,
		LastName:      lastName,
		Email:         email,
		Password:      string(hashedPassword),
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

// SignInService authenticates a user by email and password and returns the user and the error if any occurred
func SignInService(email, password string) (*models.User, error) {
	// Retrieve the user by email
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	// Compare the password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("invalid email or password")
	}

	return user, nil
}

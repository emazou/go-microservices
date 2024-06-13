package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"user-service/models"
	"user-service/repositories"
)

// SignUpService creates a new user in the database and returns the user and the error if any occurred
func SignUpService(userIn models.SignUpUserInput) (*models.User, error) {
	// Define the email regex pattern to validate the email
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched := regexp.MustCompile(emailRegex).MatchString(userIn.Email); !matched {
		return nil, errors.New("invalid email address")
	}
	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userIn.Password), bcrypt.DefaultCost)
	// Create a new user
	user := &models.User{
		Name:          userIn.Name,
		LastName:      userIn.LastName,
		Email:         userIn.Email,
		Password:      string(hashedPassword),
		Address:       userIn.Address,
		PaymentMethod: models.PaymentMethod(userIn.PaymentMethod),
		Role:          models.Role(userIn.Role),
	}
	// Create the user
	err := repositories.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return repositories.GetUserByEmail(user.Email)
}

// SignInService authenticates a user by email and password and returns the user and the error if any occurred
func SignInService(userIn models.SignInUserInput) (*models.User, error) {
	// Retrieve the user by email
	user, err := repositories.GetUserByEmail(userIn.Email)
	if err != nil {
		return user, err
	}
	// Compare the password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userIn.Password))
	return user, err
}

// GetAllUsersService retrieves all users from the database and returns the users and the error if any occurred
func GetAllUsersService() ([]models.User, error) {
	return repositories.GetAllUsers()
}

// DeleteUseByIDService deletes a user from the database by its ID and returns the error if any occurred
func DeleteUseByIDService(id string) error {
	return repositories.DeleteUserByID(id)
}

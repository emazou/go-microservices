package services

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
	"user-service/config"
	"user-service/models"
)

// Mock the CreateUser and GetUserByEmail functions in the repositories package
var (
	mockCreateUser = func(user *models.User) error {
		user.ID = "mock-id"
		return nil
	}

	mockGetUserByEmail = func(email string) (*models.User, error) {
		return &models.User{
			ID:            "mock-id",
			Name:          "John",
			LastName:      "Doe",
			Email:         email,
			Password:      "password",
			Address:       "123 Main St",
			PaymentMethod: models.CreditCard,
			Role:          models.Buyer,
			Verified:      false,
		}, nil
	}
)

func TestMain(m *testing.M) {
	config.ConnectTestDatabase()
	code := m.Run()
	config.TeardownTestDatabase()
	os.Exit(code)
}

func TestSignUpService(t *testing.T) {
	// Test the SignUpService function
	userIn := models.SignUpUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "testingSignUp@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Role:          models.Buyer,
	}

	user, err := SignUpService(userIn)
	if err != nil {
		t.Errorf("An error occurred: %v", err)
	}
	// Hash the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userIn.Password))
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
	assert.Equal(t, userIn.Name, user.Name, "User name should be the same")
	assert.Equal(t, userIn.LastName, user.LastName, "User last name should be the same")
	assert.Equal(t, userIn.Email, user.Email, "User email should be the same")
	assert.Nil(t, err, "Password should be hashed")
	assert.Equal(t, userIn.Address, user.Address, "User address should be the same")
	assert.Equal(t, userIn.PaymentMethod, user.PaymentMethod, "User payment method should be the same")
	assert.Equal(t, userIn.Role, user.Role, "User role should be the same")
}

func TestSignUpServiceInvalidEmail(t *testing.T) {
	userIn := models.SignUpUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "invalidEmail",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Role:          models.Buyer,
	}
	user, err := SignUpService(userIn)
	assert.Nil(t, user, "User should be nil")
	assert.Error(t, err, "An error should occur")
}

func TestSignUpServiceEmailAlreadyExists(t *testing.T) {
	userIn := models.SignUpUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "testingSignUp@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Role:          models.Buyer,
	}
	user, err := SignUpService(userIn)
	assert.Nil(t, user, "User should be nil")
	assert.Error(t, err, "An error should occur")
}
func TestSignInService(t *testing.T) {
	userIn := models.SignInUserInput{
		Email:    "testingSignUp@gmail.com",
		Password: "password",
	}
	user, err := SignInService(userIn)
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
	assert.Equal(t, userIn.Email, user.Email, "User email should be the same")
	assert.Nil(t, err, "No error should occur")
}

func TestSignInServiceInvalidEmail(t *testing.T) {
	userIn := models.SignInUserInput{
		Email:    "test@gmail.com",
		Password: "password",
	}
	user, err := SignInService(userIn)
	assert.Empty(t, user, "User should be empty")
	assert.Error(t, err, "An error should occur")
}

func TestSignInServiceInvalidCredentials(t *testing.T) {
	userIn := models.SignInUserInput{
		Email:    "test@gmail.com",
		Password: "invalidPassword",
	}
	user, err := SignInService(userIn)
	assert.Empty(t, user, "User should be empty")
	assert.Error(t, err, "An error should occur")
}

func TestGetAllUsersService(t *testing.T) {
	users, err := GetAllUsersService()
	assert.NotEmpty(t, users, "Users list should not be empty")
	assert.Nil(t, err, "No error should occur")
}

func TestDeleteUserByIDService(t *testing.T) {
	users, err := GetAllUsersService()
	if err != nil {
		t.Errorf("An error occurred: %v", err)
	}
	err = DeleteUserByIDService(users[0].ID)
	assert.Nil(t, err, "No error should occur")
}

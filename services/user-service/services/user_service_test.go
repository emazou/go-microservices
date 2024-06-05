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
	name := "John"
	lastName := "Doe"
	email := "testingSignUp@gmail.com"
	password := "password"
	address := "123 Main St"
	paymentMethod := models.CreditCard
	role := models.Buyer

	user, err := SignUpService(name, lastName, email, password, address, string(paymentMethod), string(role))
	if err != nil {
		t.Errorf("An error occurred: %v", err)
	}
	// Hash the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
	assert.Equal(t, name, user.Name, "User name should be the same")
	assert.Equal(t, lastName, user.LastName, "User last name should be the same")
	assert.Equal(t, email, user.Email, "User email should be the same")
	assert.Nil(t, err, "Password should be hashed")
	assert.Equal(t, address, user.Address, "User address should be the same")
	assert.Equal(t, paymentMethod, user.PaymentMethod, "User payment method should be the same")
	assert.Equal(t, role, user.Role, "User role should be the same")
}

func TestSignUpServiceInvalidEmail(t *testing.T) {
	name := "John"
	lastName := "Doe"
	email := "invalidEmail"
	password := "password"
	address := "123 Main St"
	paymentMethod := models.CreditCard
	role := models.Buyer
	user, err := SignUpService(name, lastName, email, password, address, string(paymentMethod), string(role))
	assert.Nil(t, user, "User should be nil")
	assert.Error(t, err, "An error should occur")
}

func TestSignInService(t *testing.T) {
	email := "testingSignUp@gmail.com"
	password := "password"
	user, err := SignInService(email, password)
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
	assert.Equal(t, email, user.Email, "User email should be the same")
	assert.Nil(t, err, "No error should occur")
}

func TestSignInServiceInvalidEmail(t *testing.T) {
	email := "test@gmail.com"
	password := "password"
	user, err := SignInService(email, password)
	assert.Empty(t, user, "User should be empty")
	assert.Error(t, err, "An error should occur")
}

func TestSiginServiceInvalidCredentials(t *testing.T) {
	email := "test@gmail.com"
	password := "invalidPassword"
	user, err := SignInService(email, password)
	assert.Empty(t, user, "User should be empty")
	assert.Error(t, err, "An error should occur")
}

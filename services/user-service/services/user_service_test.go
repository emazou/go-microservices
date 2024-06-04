package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"user-service/config"
	"user-service/models"
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
	assert.NotEmpty(t, user.ID, "User ID should not be empty")
	assert.Equal(t, name, user.Name, "User name should be the same")
	assert.Equal(t, lastName, user.LastName, "User last name should be the same")
	assert.Equal(t, email, user.Email, "User email should be the same")
	assert.Equal(t, password, user.Password, "User password should be the same")
	assert.Equal(t, address, user.Address, "User address should be the same")
	assert.Equal(t, paymentMethod, user.PaymentMethod, "User payment method should be the same")
	assert.Equal(t, role, user.Role, "User role should be the same")
}

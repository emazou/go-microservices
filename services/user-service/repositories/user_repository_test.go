package repositories

import (
	"github.com/stretchr/testify/assert"
	"log"
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

func TestCreateUser(t *testing.T) {
	// Test the CreateUser function
	user := &models.User{
		Name:          "John",
		LastName:      "Doe",
		Email:         "email@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Verified:      false,
		Role:          models.Buyer,
	}
	err := CreateUser(user)
	assert.NoError(t, err)
	assert.NotZerof(t, user.ID, "User ID should not be zero")
}

func TestGetUserByID(t *testing.T) {
	// Test the GetUserByID function
	user := &models.User{
		Name:          "John",
		LastName:      "Doe",
		Email:         "email2@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Verified:      false,
		Role:          models.Buyer,
	}
	err := CreateUser(user)
	assert.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	// Test the GetUserByEmail function
	user := &models.User{
		Name:          "John",
		LastName:      "Doe",
		Email:         "email3@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: models.CreditCard,
		Verified:      false,
		Role:          models.Buyer,
	}
	err := CreateUser(user)
	assert.NoError(t, err)

	userByEmail, err := GetUserByEmail(user.Email)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, userByEmail.ID)
}

func TestGetAllUsers(t *testing.T) {
	// Test the GetAllUsers function
	users, err := GetAllUsers()
	assert.NoError(t, err)
	assert.NotZerof(t, len(users), "Users list should not be empty")
}

func TestUpdateUser(t *testing.T) {
	users, err := GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	myUserID := users[0].ID
	updatedUser := &models.User{
		ID:    myUserID,
		Email: "emailupdated@gmail.com",
	}
	err = UpdateUser(updatedUser)
	updatedUserByID, err := GetUserByID(myUserID)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Email, updatedUserByID.Email)
}

func TestDeleteUserByID(t *testing.T) {
	users, err := GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	myUserID := users[0].ID
	err = DeleteUserByID(myUserID)
	assert.NoError(t, err)
	deletedUser, err := GetUserByID(myUserID)
	assert.Error(t, err)
	assert.Empty(t, deletedUser)
}

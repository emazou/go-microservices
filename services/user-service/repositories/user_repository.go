package repositories

import (
	"user-service/config"
	"user-service/models"
)

// CreateUser creates a new user in the database and returns the error if any occurred
func CreateUser(user *models.User) error {
	// Create a new user
	return config.DB.Create(&user).Error
}

// GetUserByID retrieves a user from the database by its ID and returns the error if any occurred
func GetUserByID(id string) (*models.User, error) {
	// Find a user by its ID
	var user *models.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

// GetUserByEmail retrieves a user from the database by its email and returns the error if any occurred
func GetUserByEmail(email string) (*models.User, error) {
	// Find a user by its email
	var user *models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// UpdateUser updates a user in the database and returns the error if any occurred
func UpdateUser(user *models.User) error {
	// Update a user
	return config.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

// DeleteUserByID deletes a user from the database by its ID and returns the error if any occurred
func DeleteUserByID(id string) error {
	// Check if the user exists
	var user models.User
	err := config.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return err
	}
	// Delete a user by its ID
	return config.DB.Where("id = ?", id).Delete(&models.User{}).Error
}

// GetAllUsers retrieves all users from the database and returns the error if any occurred
func GetAllUsers() ([]models.User, error) {
	// Find all users
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

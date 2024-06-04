package models

import (
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestUserModel(t *testing.T) {
	// Create a new user
	user := &User{
		Name:          "John",
		LastName:      "Doe",
		Email:         "email@gmail.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: CreditCard,
		Verified:      false,
		Role:          Buyer,
	}
	// Verify the user fields
	tests.AssertEqual(t, user.Name, "John")
	tests.AssertEqual(t, user.LastName, "Doe")
	tests.AssertEqual(t, user.Email, "email@gmail.com")
	tests.AssertEqual(t, user.Password, "password")
	tests.AssertEqual(t, user.Address, "123 Main St")
	tests.AssertEqual(t, user.PaymentMethod, CreditCard)
	tests.AssertEqual(t, user.Verified, false)
	tests.AssertEqual(t, user.Role, Buyer)
}

func TestPaymentMethodEnum(t *testing.T) {
	// Verify the payment method values
	tests.AssertEqual(t, CreditCard, "CREDIT_CARD")
	tests.AssertEqual(t, DebitCard, "DEBIT_CARD")
}

func TestRoleEnum(t *testing.T) {
	// Verify the role values
	tests.AssertEqual(t, Seller, "SELLER")
	tests.AssertEqual(t, Admin, "ADMIN")
	tests.AssertEqual(t, Buyer, "BUYER")
}

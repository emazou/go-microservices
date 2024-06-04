package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentMethod string
type Role string

const (
	CreditCard PaymentMethod = "CREDIT_CARD"
	DebitCard  PaymentMethod = "DEBIT_CARD"
)

const (
	Seller Role = "SELLER"
	Admin  Role = "ADMIN"
	Buyer  Role = "BUYER"
)

type User struct {
	ID            string        `gorm:"type:varchar(36);primary_key" json:"id"`
	Name          string        `gorm:"type:varchar(50);not null" json:"name"`
	LastName      string        `gorm:"type:varchar(50);not null" json:"last_name"`
	Email         string        `gorm:"type:varchar(80);not null; unique" json:"email"`
	Password      string        `gorm:"type:varchar(255);not null" json:"password"`
	Address       string        `gorm:"type:varchar(50);not null" json:"address"`
	PaymentMethod PaymentMethod `gorm:"type:enum('CREDIT_CARD','DEBIT_CARD','PAYPAL','BANK_TRANSFER');not null" json:"payment_method"`
	Verified      bool          `gorm:"type:boolean;default:false" json:"verified"`
	Role          Role          `gorm:"type:enum('SELLER','ADMIN','BUYER');not null" json:"role"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Generate a new UUID
	random, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// Set the UUID as the user ID
	user.ID = random.String()
	return nil

}

package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewJWTService(t *testing.T) {
	jwtService := NewJWTService()
	assert.NotNil(t, jwtService, "JWT service should not be nil")
}

func TestJwtServiceGenerateToken(t *testing.T) {
	id := "mock-id"
	email := "mock-email"
	jwtService := NewJWTService()
	token, err := jwtService.GenerateToken(id, email)
	assert.NoError(t, err, "An error occurred")
	assert.NotEmpty(t, token, "Token should not be empty")
}

func TestJwtServiceValidateToken(t *testing.T) {
	id := "mock-id"
	email := "mock-email"
	jwtService := NewJWTService()
	token, err := jwtService.GenerateToken(id, email)
	assert.NoError(t, err, "An error occurred")
	assert.NotEmpty(t, token, "Token should not be empty")
	_, err = jwtService.ValidateToken(token)
	assert.NoError(t, err, "An error occurred")
}

func TestJwtServiceValidateTokenInvalidToken(t *testing.T) {
	jwtService := NewJWTService()
	_, err := jwtService.ValidateToken("invalid-token")
	assert.Error(t, err, "An error should occur")
}

func TestDecodeToken(t *testing.T) {

}

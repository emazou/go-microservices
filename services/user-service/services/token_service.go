package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"time"
)

var jwtKey []byte

func init() {
	godotenv.Load(filepath.Join("..", ".env"))
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}

type JWTService interface {
	GenerateToken(id, email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	DecodeToken(token string) (*Claims, error)
}

type jwtService struct{}

func NewJWTService() JWTService {
	return &jwtService{}
}

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(id, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func (j jwtService) ValidateToken(token string) (*jwt.Token, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return tkn, nil
}

func (j *jwtService) DecodeToken(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

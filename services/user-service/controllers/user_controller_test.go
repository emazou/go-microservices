package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"user-service/config"
	"user-service/models"
)

const (
	BaseUrl   = "/api/v1/users"
	SignupUrl = "/signup"
	SigninUrl = "/signin"
)

func TestMain(m *testing.M) {
	config.ConnectTestDatabase()
	code := m.Run()
	config.TeardownTestDatabase()
	os.Exit(code)
}

func setupTestRouter() *gin.Engine {
	// Set up the routes
	router := gin.Default()
	v1 := router.Group(BaseUrl)
	{
		v1.POST(SignupUrl, SignUp)
		v1.POST(SigninUrl, SignIn)
	}
	return router
}

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupTestRouter()

	input := models.SignUpUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "john.doe@example.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: string(models.CreditCard),
		Role:          string(models.Buyer),
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SignupUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	data := response["data"].(map[string]interface{})
	assert.Equal(t, input.Name, data["name"])
	assert.Equal(t, input.LastName, data["last_name"])
	assert.Equal(t, input.Email, data["email"])
	assert.Equal(t, input.Address, data["address"])
	assert.Equal(t, input.PaymentMethod, data["payment_method"])
	assert.Equal(t, input.Role, data["role"])
}

func TestSignUpInvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupTestRouter()

	input := `{"name": "Invalid name"}`
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SignupUrl, bytes.NewBuffer([]byte(input)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignUpErrorService(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()

	input := models.SignUpUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "john.doe@example.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: string(models.CreditCard),
		Role:          string(models.Buyer),
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SignupUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSignIn(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	input := models.SignUpUserInput{
		Email:    "john.doe@example.com",
		Password: "password",
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SigninUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInInvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	input := `{}`
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SigninUrl, bytes.NewBuffer([]byte(input)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignInInvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	input := models.SignInUserInput{
		Email:    "test@gmail.com",
		Password: "password",
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, BaseUrl+SigninUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

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
	"user-service/middlewares"
	"user-service/models"
	"user-service/services"
	"user-service/utils"
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
	jwtService := services.NewJWTService()
	v1 := router.Group(utils.BaseUrl)
	{
		v1.POST(utils.SignupUrl, SignUp)
		v1.POST(utils.SigninUrl, SignIn)
		v1.DELETE("/:id", middlewares.AuthMiddleware(jwtService), DeleteUserByID)
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
		PaymentMethod: models.CreditCard,
		Role:          models.Buyer,
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SignupUrl, bytes.NewBuffer(jsonInput))
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
	assert.Equal(t, string(input.PaymentMethod), data["payment_method"])
	assert.Equal(t, string(input.Role), data["role"])
}

func TestSignUpInvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupTestRouter()

	input := `{"name": "Invalid name"}`
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SignupUrl, bytes.NewBuffer([]byte(input)))
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
		PaymentMethod: models.CreditCard,
		Role:          models.Buyer,
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SignupUrl, bytes.NewBuffer(jsonInput))
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
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SigninUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInInvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	input := `{}`
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SigninUrl, bytes.NewBuffer([]byte(input)))
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
	req, _ := http.NewRequest(http.MethodPost, utils.BaseUrl+utils.SigninUrl, bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeleteUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	users, err := services.GetAllUsersService()
	if err != nil {
		t.Errorf("An error occurred: %v", err)
	}
	user := users[0]
	jwtService := services.NewJWTService()
	token, _ := jwtService.GenerateToken(user.ID, user.Email)
	req, _ := http.NewRequest(http.MethodDelete, utils.BaseUrl+"/"+user.ID, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUserByIDWithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	jwtService := services.NewJWTService()
	token, _ := jwtService.GenerateToken("13", "test@gmail.com")
	req, _ := http.NewRequest(http.MethodDelete, utils.BaseUrl+"/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestDeleteUserByIDWithTokenInvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := setupTestRouter()
	jwtService := services.NewJWTService()
	token, _ := jwtService.GenerateToken("1", "test@gmail.com")
	req, _ := http.NewRequest(http.MethodDelete, utils.BaseUrl+"/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

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

func TestMain(m *testing.M) {
	config.ConnectTestDatabase()
	code := m.Run()
	config.TeardownTestDatabase()
	os.Exit(code)
}

func setupTestRouter() *gin.Engine {
	// Set up the routes
	router := gin.Default()
	v1 := router.Group("/api/v1/users")
	{
		v1.POST("/signup", SignUp)
	}
	return router
}

func TestSignUp(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupTestRouter()

	input := CreateUserInput{
		Name:          "John",
		LastName:      "Doe",
		Email:         "john.doe@example.com",
		Password:      "password",
		Address:       "123 Main St",
		PaymentMethod: string(models.CreditCard),
		Role:          string(models.Buyer),
	}
	jsonInput, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/signup", bytes.NewBuffer(jsonInput))
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

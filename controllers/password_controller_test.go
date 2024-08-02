package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the Password Service
type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) CalculateSteps(password string) int {
	args := m.Called(password)
	return args.Int(0)
}

func TestPasswordStrengthHandler(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Set up the mock service
	mockService := new(MockPasswordService)
	mockService.On("CalculateSteps", "aA1").Return(3)
	mockService.On("CalculateSteps", "1445D1cd").Return(0)

	// Inject the mock service into the controller
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		PasswordStrengthHandler(c)
	})

	tests := []struct {
		requestBody          string
		expectedStatusCode   int
		expectedResponseBody map[string]interface{}
	}{
		{
			requestBody:          `{"init_password": "aA1"}`,
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: map[string]interface{}{"num_of_steps": 3},
		},
		{
			requestBody:          `{"init_password": "1445D1cd"}`,
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: map[string]interface{}{"num_of_steps": 0},
		},
		{
			requestBody:          `{"init_password": "invalid"}`,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: map[string]interface{}{"error": "Invalid input"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.requestBody, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/api/strong_password_steps", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			var responseBody map[string]interface{}
			_ = json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.Equal(t, tt.expectedResponseBody, responseBody)
		})
	}
}

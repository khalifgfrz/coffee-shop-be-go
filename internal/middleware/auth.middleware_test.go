package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedBody   string
		roles          []string
	}{
		{
			name:           "Missing Authorization Header",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"message":"Unauthorized"}`,
		},
		{
			name:           "Invalid Bearer Token",
			token:          "InvalidBearerToken",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"message":"Invalid Bearer Token"}`,
		},
		{
			name:           "Valid Token with Allowed Role",
			token:          "validToken",
			expectedStatus: http.StatusOK,
			expectedBody:   "Success",
			roles:          []string{"admin"},
		},
		{
			name:           "Valid Token with Forbidden Role",
			token:          "validToken",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"message":"Forbidden: Insufficient permissions"}`,
			roles:          []string{"user"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			router.Use(MockAuthMiddleware(tt.roles...))

			router.GET("/test", func(c *gin.Context) {
				c.String(http.StatusOK, "Success")
			})

			req, _ := http.NewRequest("GET", "/test", nil)
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

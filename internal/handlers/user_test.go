package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"task-manager-app/internal/services"
	"task-manager-app/internal/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	userService := services.Service{}
	handler := Handler(&userService)

	tests := []struct {
		Name       string
		Mock       []byte
		StatusCode int
	}{
		{
			Name:       "success",
			Mock:       utils.OpenMock("../mocks/user_mocks/create_user.json"),
			StatusCode: http.StatusCreated,
		},
		{
			Name:       "error",
			Mock:       utils.OpenMock("../mocks/user_mocks/create_error.json"),
			StatusCode: http.StatusInternalServerError,
		},
		{
			Name:       "empty param",
			Mock:       utils.OpenMock("../mocks/user_mocks/empty_param.json"),
			StatusCode: http.StatusBadRequest,
		},
		{
			Name:       "left param",
			Mock:       []byte{},
			StatusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.Mock))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			router := gin.Default()
			router.POST("/users", handler.Create)

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.StatusCode, rec.Code)

		})
	}
}

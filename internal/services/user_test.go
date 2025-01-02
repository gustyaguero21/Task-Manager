package services

import (
	"context"
	"errors"
	"task-manager-app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	service := Service{}

	tests := []struct {
		Name          string
		User          models.User
		ExpectedError error
	}{
		{
			Name:          "Success",
			User:          models.User{Name: "John", Surname: "Doe", Email: "johndoe@example.com", Password: "Password1234"},
			ExpectedError: nil,
		},
		{
			Name:          "Invalid email",
			User:          models.User{Name: "John", Surname: "Doe", Email: "johndoeexample.com", Password: "Password1234"},
			ExpectedError: errors.New("invalid email address"),
		},
		{
			Name:          "Invalid password",
			User:          models.User{Name: "John", Surname: "Doe", Email: "johndoe@example.com", Password: "password1234"},
			ExpectedError: errors.New("invalid password"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := service.CreateUser(ctx, tt.User)

			assert.Equal(t, tt.ExpectedError, err)
		})
	}
}

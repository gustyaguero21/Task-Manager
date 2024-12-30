package services

import (
	"context"
	"task-manager-app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Ok(t *testing.T) {
	//given
	ctx := context.Background()
	service := &Service{}
	user := models.User{
		Name:     "John",
		Surname:  "Doe",
		Email:    "johndoe@example.com",
		Password: "Password1234",
	}

	//act

	created, err := service.CreateUser(ctx, user)

	//assert

	assert.Nil(t, err)
	assert.NotNil(t, created)
}

func TestInvalid_Email(t *testing.T) {
	//given
	ctx := context.Background()
	service := &Service{}
	user := models.User{
		Name:     "John",
		Surname:  "Doe",
		Email:    "johndoe",
		Password: "Password1234",
	}

	//act
	_, err := service.CreateUser(ctx, user)

	//assert
	assert.Error(t, err)

}
func TestInvalid_Password(t *testing.T) {
	//given
	ctx := context.Background()
	service := &Service{}
	user := models.User{
		Name:     "John",
		Surname:  "Doe",
		Email:    "johndoe@example.com",
		Password: "password1234",
	}

	//act
	_, err := service.CreateUser(ctx, user)

	//assert
	assert.Error(t, err)

}

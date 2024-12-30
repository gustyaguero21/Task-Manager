package service

import (
	"task-manager-app/internal/models"
	"task-manager-app/internal/service"
	"testing"
)

func TestCreateUser_Ok(t *testing.T) {
	//given

	user := models.User{
		Name:     "John",
		Surname:  "Doe",
		Email:    "johndoe@example.com",
		Password: "Password1234",
	}

	//act

	created, err := 
}

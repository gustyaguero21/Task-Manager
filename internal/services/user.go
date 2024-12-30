package services

import (
	"context"
	"errors"
	"task-manager-app/internal/models"

	"github.com/google/uuid"
	"github.com/gustyaguero21/Go-toolkit/pkg/encrypter"
	"github.com/gustyaguero21/Go-toolkit/pkg/validator"
)

type UserService struct{}

func (us *UserService) CreateUser(ctx context.Context, user models.User) (createdUser models.User, err error) {

	if !validator.ValidateEmail(user.Email) {
		return models.User{}, errors.New("invalid email address")
	}

	if !validator.ValidatePassword(user.Password) {
		return models.User{}, errors.New("invalid password")
	}

	hashedPwd, hashErr := encrypter.PasswordEncrypter(user.Password)
	if hashErr != nil {
		return models.User{}, errors.New("error creating password hash")
	}

	uID := uuid.New()

	newUser := models.User{
		ID:       uID.String(),
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: string(hashedPwd),
	}

	return newUser, nil
}

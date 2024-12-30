package services

import (
	"context"
	"task-manager-app/internal/models"
)

type UserServices interface {
	CreateUser(ctx context.Context, user models.User) (createdUser models.User, err error)
}

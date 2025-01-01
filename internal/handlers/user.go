package handlers

import (
	"net/http"
	"task-manager-app/internal/models"
	"task-manager-app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/gustyaguero21/Go-toolkit/pkg/web"
)

type UserHandler struct {
	userService services.UserServices
}

func Handler(userService services.UserServices) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Create(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		web.NewError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if user.Name == "" || user.Surname == "" || user.Email == "" || user.Password == "" {
		web.NewError(ctx, http.StatusBadRequest, "all fields are required")
		return
	}

	created, createdErr := uh.userService.CreateUser(ctx, user)
	if createdErr != nil {
		web.NewError(ctx, http.StatusInternalServerError, "error creating user "+createdErr.Error())
		return
	}

	ctx.JSON(http.StatusCreated, createUserResponse(http.StatusOK, "success", created))
}

func createUserResponse(status int, message string, created models.User) *models.CreateUserResponse {
	return &models.CreateUserResponse{
		StatusCode:  status,
		Message:     message,
		CreatedUser: created,
	}
}

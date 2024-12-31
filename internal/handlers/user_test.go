package handlers

import (
	"fmt"
	"task-manager-app/internal/models"
	"task-manager-app/internal/utils"
	"testing"
)

func TestCreateUser(t *testing.T) {

	body, _ := utils.OpenMock[models.User]("../mocks/create_user.json")
	fmt.Println(body)

}

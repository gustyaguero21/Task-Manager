package router

import (
	"task-manager-app/internal/handlers"
	"task-manager-app/internal/services"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {
	api := r.Group("/api/task-manager")

	userService := services.Service{}
	handler := handlers.Handler(&userService)

	api.POST("/create", handler.Create)

}

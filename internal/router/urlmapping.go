package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {
	api := r.Group("/api/task-manager")

	api.POST("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})
}

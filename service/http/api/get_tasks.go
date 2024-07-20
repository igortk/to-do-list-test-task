package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"to-do-list-test-task/dto"
)

func GetTasksHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tasks []dto.Task
		db.Find(&tasks)

		ctx.JSON(http.StatusOK, dto.GetTasksResponse{Task: tasks})
	}
}

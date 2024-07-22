package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do-list-test-task/dto"
	"to-do-list-test-task/storage/postgre"
)

func GetTasksHandler(pClient *postgre.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tasks := pClient.GetAllTasks()

		ctx.JSON(http.StatusOK, dto.GetTasksResponse{Task: tasks})
	}
}

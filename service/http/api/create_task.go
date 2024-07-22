package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"time"
	"to-do-list-test-task/dto"
	"to-do-list-test-task/storage/postgre"
)

func CreateTaskHandler(pClient *postgre.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := http.StatusOK

		tsk, err := process(ctx)
		if err != nil {
			status = http.StatusBadRequest
		}

		pClient.InsertTask(tsk)
		ctx.JSON(status, dto.CreateTaskResponse{Task: tsk})
	}
}

func process(ctx *gin.Context) (*dto.Task, error) {
	req := &dto.CreateTaskRequest{}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return nil, err
	}

	tsk := &dto.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return tsk, nil
}

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"io"
	"net/http"
	"time"
	"to-do-list-test-task/dto"
)

func UpdateTaskByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tsk dto.Task
		req := &dto.UpdateTaskByIdRequest{}
		status := http.StatusBadRequest

		err := process2(ctx, req)
		if err == nil {

			id, _ := ctx.Params.Get(ReqTaskId)

			db.Model(&dto.Task{}).Where("id = ?", id).
				Update("title", req.Title).
				Update("description", req.Description).
				Update("due_date", req.DueDate).
				Update("updated_at", time.Now().Format("2006-01-02 15:04:05"))

			db.Find(&tsk, id)

			status = http.StatusOK
		}

		ctx.JSON(status, &dto.UpdateTaskByIdResponse{Task: tsk})
	}
}

func process2(ctx *gin.Context, req *dto.UpdateTaskByIdRequest) error {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}

	return nil
}

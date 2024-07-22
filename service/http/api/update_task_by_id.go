package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"net/http"
	"to-do-list-test-task/dto"
	"to-do-list-test-task/storage/postgre"
)

func UpdateTaskByIdHandler(pClient *postgre.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tsk dto.Task
		req := &dto.UpdateTaskByIdRequest{}
		status := http.StatusBadRequest

		err := process2(ctx, req)
		if err == nil {

			id, _ := ctx.Params.Get(ReqTaskId)

			tsk = pClient.UpdateTaskById(id, req)

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

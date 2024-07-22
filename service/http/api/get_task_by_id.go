package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"to-do-list-test-task/dto"
	"to-do-list-test-task/storage/postgre"
)

const ReqTaskId = "id"

func GetTaskByIdHandler(pClient *postgre.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		status := http.StatusOK
		id, _ := ctx.Params.Get(ReqTaskId)

		tsk := pClient.GetTaskById(id)
		if reflect.DeepEqual(tsk, dto.Task{}) {
			status = http.StatusNotFound
		}

		ctx.JSON(status, &dto.GetTaskByIdResponse{Task: tsk})
	}
}

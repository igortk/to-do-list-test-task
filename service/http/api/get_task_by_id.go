package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"reflect"
	"to-do-list-test-task/dto"
)

const ReqTaskId = "id"

func GetTaskByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tsk dto.Task

		status := http.StatusOK
		id, _ := ctx.Params.Get(ReqTaskId)

		db.Find(&tsk, id)
		if reflect.DeepEqual(tsk, dto.Task{}) {
			status = http.StatusNotFound
		}

		ctx.JSON(status, &dto.GetTaskByIdResponse{Task: tsk})
	}
}

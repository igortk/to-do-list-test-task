package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"to-do-list-test-task/dto"
)

func DeleteTaskByIdHandler(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get(ReqTaskId)

		db.Delete(&dto.Task{}, id)

		ctx.JSON(http.StatusNoContent, "")
	}
}

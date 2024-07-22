package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do-list-test-task/storage/postgre"
)

func DeleteTaskByIdHandler(pClient *postgre.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get(ReqTaskId)

		pClient.DeleteTaskById(id)

		ctx.JSON(http.StatusNoContent, "")
	}
}

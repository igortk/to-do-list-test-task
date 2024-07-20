package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"to-do-list-test-task/dto"
)

func main() {
	e := gin.Default()

	e.POST("/tasks", HandlerCreateTask)

	err := e.Run(":8888")
	if err != nil {
		log.Fatal(err)
	}
}

func HandlerCreateTask(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("D:\\to-do-list-test-task\\to-do-list.db"), &gorm.Config{})
	if err != nil {
		log.Error(err)
	}

	err = db.AutoMigrate(&dto.Task{})
	if err != nil {
		log.Error()
	}

	req := &dto.CreateTaskRequest{}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		log.Error(err)
	}

	resp := &dto.CreateTaskResponse{
		Id:          0,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		CreatedAt:   req.DueDate + "1",
		UpdatedAt:   req.DueDate + "2",
	}
	ctx.JSON(http.StatusOK, resp)
}

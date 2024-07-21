package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"to-do-list-test-task/config"
	"to-do-list-test-task/dto"
	"to-do-list-test-task/service/http/api"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	e := gin.Default()
	db, err := gorm.Open(postgres.Open(cfg.PostgresConfig.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&dto.Task{})
	if err != nil {
		log.Error(err)
	}

	opAdGroup := e.Group("/")

	opAdGroup.POST("/tasks", api.CreateTaskHandler(db))
	opAdGroup.GET("/tasks", api.GetTasksHandler(db))
	opAdGroup.GET("/tasks/:id", api.GetTaskByIdHandler(db))
	opAdGroup.PUT("/tasks/:id", api.UpdateTaskByIdHandler(db))

	err = e.Run(cfg.HttpConfig.Port)
	if err != nil {
		log.Fatal(err)
	}
}

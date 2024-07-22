package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
	"to-do-list-test-task/config"
	"to-do-list-test-task/dto"
)

type Client struct {
	DB *gorm.DB
}

func New(cfg *config.PostgresConfig) (*Client, error) {
	db, err := gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&dto.Task{})
	if err != nil {
		return nil, err
	}

	return &Client{
		DB: db,
	}, nil
}

func (c *Client) InsertTask(tsk *dto.Task) {
	c.DB.Create(tsk)
}

func (c *Client) DeleteTaskById(id string) {
	c.DB.Delete(&dto.Task{}, id)
}

func (c *Client) GetTaskById(id string) dto.Task {
	var tsk dto.Task
	c.DB.Find(&tsk, id)

	return tsk
}

func (c *Client) GetAllTasks() []dto.Task {
	var tasks []dto.Task
	c.DB.Find(&tasks)

	return tasks
}

func (c *Client) UpdateTaskById(id string, req *dto.UpdateTaskByIdRequest) dto.Task {
	var tsk dto.Task

	c.DB.Model(&dto.Task{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       req.Title,
		"description": req.Description,
		"due_date":    req.DueDate,
		"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
	})

	c.DB.Find(&tsk, id)

	return tsk
}

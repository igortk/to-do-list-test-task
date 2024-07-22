package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

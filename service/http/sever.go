package http

import (
	"github.com/gin-gonic/gin"
	"to-do-list-test-task/config"
	"to-do-list-test-task/service/http/api"
	"to-do-list-test-task/storage/postgre"
)

type Server struct {
	e       *gin.Engine
	pClient *postgre.Client
	cfg     *config.HttpConfig
}

func New(cfg *config.HttpConfig, pClient *postgre.Client) *Server {
	return &Server{
		cfg:     cfg,
		pClient: pClient,
	}
}

func (s *Server) initEngine() {
	e := gin.Default()
	opAdGroup := e.Group("/")

	opAdGroup.POST("/tasks", api.CreateTaskHandler(s.pClient.DB))
	opAdGroup.GET("/tasks", api.GetTasksHandler(s.pClient.DB))
	opAdGroup.GET("/tasks/:id", api.GetTaskByIdHandler(s.pClient.DB))
	opAdGroup.PUT("/tasks/:id", api.UpdateTaskByIdHandler(s.pClient.DB))
	opAdGroup.DELETE("/tasks/:id", api.DeleteTaskByIdHandler(s.pClient.DB))

	s.e = e
}

func (s *Server) Run() error {
	s.initEngine()

	err := s.e.Run(s.cfg.Port)

	return err
}

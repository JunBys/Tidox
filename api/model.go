package api

import (
	"fmt"
	"tidox/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	App         *gin.Engine
	StorageTask *storage.Task
}

func NewServer(st *storage.Task) *Server {
	return &Server{
		App:         gin.Default(),
		StorageTask: st,
	}
}

// Run 启动gui gin服务器
func (s *Server) Run(addr string) error {
	// 导入模板和静态文件
	s.App.LoadHTMLGlob("templates/*")
	s.App.Static("/static", "static")

	// 启动前端时间渲染服务
	s.runFront()

	// 启动后端渲染服务
	s.runBackend(s.StorageTask)

	// 启动服务
	err := s.App.Run(addr)
	if err != nil {
		return fmt.Errorf("gin服务器启动失败:%w", err)
	}
	return nil
}

package api

import (
	"tidox/service"
	"tidox/storage"

	"github.com/gin-gonic/gin"
)

func (s *Server) runBackend(storageTask *storage.Task) {
	api := s.App.Group("/backend")
	svc := service.NewAPI(storageTask)

	{
		// 获取NewTask子页
		api.GET("/newTask", func(ctx *gin.Context) {
			data := svc.GetTask("NewTask")
			ctx.HTML(200, "new_task.html", gin.H{"data": data})
		})

		// 获取NowTask子页面
		api.GET("/nowTask", func(ctx *gin.Context) {
			data := svc.GetTask("NowTask")
			ctx.HTML(200, "now_task.html", gin.H{"data": data})
		})

		// 获取OldTask子页面
		api.GET("/oldTask", func(ctx *gin.Context) {
			data := svc.GetTask("OldTask")
			ctx.HTML(200, "old_task.html", gin.H{"data": data})
		})
	}

}

package api

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) runFront() {
	api := s.App.Group("/front")
	{
		// 获取主窗口
		api.GET("/index", func(ctx *gin.Context) {
			ctx.HTML(200, "manager.html", gin.H{})
		})

		// 获取左侧panel，通过post表单决定哪个窗格被激活
		// 会将表单的所有数据附加渲染到模版
		api.POST("/leftPanel", func(ctx *gin.Context) {
			//解析表单数据
			err := ctx.Request.ParseForm()
			if err != nil {
				ctx.String(500, "post解析失败")
			}

			//表单数据一个key可能有多个值，所以返回类型是map[string][]string
			//取第一个值
			parms := make(map[string]string)
			for k, v := range ctx.Request.Form {
				if len(v) > 0 {
					parms[k] = v[0]
				}
			}

			ctx.HTML(200, "left-panel.html", parms)
		})

	}
}

package router

import (
	"gin-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func SetupRoutes(r *gin.Engine, container do.Injector) {
	// 全局中间件

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		utils.Success(c, "api/health")
	})

	// API 路由
	api := r.Group("/api")
	ApiRouter(api, container)
}

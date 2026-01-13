package router

import (
	"gin-api/internal/api/health"

	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
)

func ApiRouter(r *gin.RouterGroup, container do.Injector) {
	h := do.MustInvoke[health.Handler](container)
	r.GET("/health", h.Health())
}

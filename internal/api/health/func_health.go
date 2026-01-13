package health

import (
	"gin-api/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *handler) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.Success(c, nil)
	}
}

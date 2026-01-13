package health

import (
	"gin-api/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/samber/do/v2"
	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	Health() gin.HandlerFunc
}
type handler struct {
	logger *zap.Logger
}

func New(i do.Injector) (Handler, error) {
	return &handler{
		logger: do.MustInvoke[*config.LoggerService](i).Logger,
	}, nil
}
func (h *handler) i() {}

package injector

import (
	"gin-api/internal/api/health"
	"gin-api/internal/config"

	"github.com/samber/do/v2"
)

func SetupInjector() do.Injector {
	injector := do.New()
	// 配置层(必须最先注册)
	do.Provide(injector, config.NewConfig)
	do.Provide(injector, config.NewLogger)
	do.Provide(injector, config.NewDB)
	do.Provide(injector, config.NewRedis)
	do.Provide(injector, config.NewQueue)

	// 注册 handlers
	do.Provide(injector, health.New)
	return injector
}

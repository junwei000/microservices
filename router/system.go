package router

import (
	"microservices/cache"
	"microservices/controller"
	"microservices/repo"
	"microservices/service"
)

func init() {
	systemController := controller.NewSystemController(repo.NewFactory(), cache.NewFactory(), service.NewFactory())
	router.GET("/system/health", systemController.Health)
}

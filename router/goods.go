package router

import (
	"microservices/cache"
	"microservices/controller"
	"microservices/repo"
	"microservices/router/middleware"
	"microservices/service"
)

func init() {
	goodsController := controller.NewGoodsController(repo.NewFactory(), cache.NewFactory(), service.NewFactory())
	router.GET("/goods", middleware.Authenticate(), goodsController.GetList)
}

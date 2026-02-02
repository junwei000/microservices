package router

import (
	"microservices/cache"
	"microservices/controller"
	"microservices/repo"
	"microservices/router/middleware"
	"microservices/service"
)

func init() {
	userController := controller.NewUserController(repo.NewFactory(), cache.NewFactory(), service.NewFactory())
	router.GET("/userinfo", middleware.Authenticate(), userController.GetUserInfo)
	router.POST("/users/:id/edit", middleware.Authenticate(), userController.Edit)
}

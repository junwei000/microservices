package router

import (
	"microservices/cache"
	"microservices/controller"
	"microservices/repo"
	"microservices/router/middleware"
	"microservices/service"
)

func init() {
	authController := controller.NewAuthController(repo.NewFactory(), cache.NewFactory(), service.NewFactory())
	router.POST("/auth/login", authController.Login)
	router.POST("/auth/logout", middleware.Authenticate(), authController.Logout)
	router.POST("/auth/change-password", middleware.Authenticate(), authController.ChangePassword)
	router.POST("/auth/change-password-by-phone", authController.ChangePasswordByPhone)
	router.POST("/auth/register", authController.Register)
}

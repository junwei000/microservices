package router

import (
	"microservices/cache"
	"microservices/controller"
	"microservices/repo"
	"microservices/service"
)

func init() {
	notifyController := controller.NewNotifyController(repo.NewFactory(), cache.NewFactory(), service.NewFactory())
	router.POST("/notify/sms", notifyController.SendSms)
}

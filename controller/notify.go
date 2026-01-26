package controller

import (
	"github.com/gin-gonic/gin"
	"microservices/cache"
	"microservices/entity/request"
	"microservices/entity/response"
	"microservices/logic"
	"microservices/model"
	"microservices/pkg/util"
	"microservices/service"
	"strconv"
)

type NotifyController interface {
	SendSms(c *gin.Context, req request.SendSms) (*response.SendSms, error)
	SendEmail(c *gin.Context) (*response.SendEmail, error)
}

type notifyController struct {
	logic logic.Factory
}

func (n *notifyController) SendSms(c *gin.Context, req request.SendSms) (*response.SendSms, error) {
	if err := n.logic.Notify().SendSmsCode(c.Request.Context(), req.Phone, strconv.Itoa(util.RandomN(4))); err != nil {
		return nil, err
	}
	return nil, nil
}

func (n *notifyController) SendEmail(c *gin.Context) (*response.SendEmail, error) {
	//TODO implement me
	panic("implement me")
}

func NewNotifyController(model model.Factory, cache cache.Factory, service service.Factory) NotifyController {
	return &notifyController{
		logic: logic.NewLogic(model, cache, service),
	}
}

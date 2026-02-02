package controller

import (
	"github.com/gin-gonic/gin"
	"microservices/cache"
	"microservices/entity/response"
	"microservices/logic"
	"microservices/repo"
	"microservices/service"
)

type SystemController interface {
	Health(c *gin.Context) (*response.Health, error)
}

type systemController struct {
	logic logic.Factory
}

// Health .
func (s *systemController) Health(c *gin.Context) (*response.Health, error) {
	return &response.Health{
		Status: "ok",
	}, nil
}

func NewSystemController(model repo.Factory, cache cache.Factory, service service.Factory) SystemController {
	return &systemController{
		logic: logic.NewLogic(model, cache, service),
	}
}

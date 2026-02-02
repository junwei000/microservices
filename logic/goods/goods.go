package goods

import (
	"context"
	"microservices/cache"
	"microservices/entity/model"
	repo "microservices/repo"
	"microservices/service"
)

type Logic interface {
	GetList(ctx context.Context) ([]*model.Goods, error)
}

type logic struct {
	repo  repo.Factory
	cache cache.Factory
	srv   service.Factory
}

func NewLogic(repo repo.Factory, cache cache.Factory, service service.Factory) Logic {
	return &logic{repo: repo, cache: cache, srv: service}
}

func (l *logic) GetList(ctx context.Context) ([]*model.Goods, error) {
	return l.repo.Goods().GetList(ctx)
}

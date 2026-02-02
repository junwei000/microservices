package notify

import (
	"context"
	"microservices/cache"
	"microservices/repo"
	"microservices/service"
)

type Logic interface {
	SendSmsCode(ctx context.Context, phone string, code string) error
}

type logic struct {
	repo  repo.Factory
	cache cache.Factory
	srv   service.Factory
}

func (n *logic) SendSmsCode(ctx context.Context, phone string, code string) error {
	if err := n.cache.Auth().SetSmsCode(ctx, phone, code); err != nil {

	}
	return n.srv.Aliyun().SendSMSCode(ctx, phone, code)
}

func NewNotify(repo repo.Factory, cache cache.Factory, service service.Factory) Logic {
	return &logic{
		repo:  repo,
		cache: cache,
		srv:   service,
	}
}

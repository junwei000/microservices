package callback

import (
	"context"
	"github.com/go-pay/gopay"
	"microservices/cache"
	entity "microservices/entity/model"
	"microservices/entity/request"
	"microservices/repo"
	"microservices/service"
)

type Logic interface {
	GoogleCallback(ctx context.Context, code string) (*entity.User, string, error)
	HandleAlipayNotify(ctx context.Context, payload map[string]string) error
	HandleAlipayCallback(ctx context.Context, params map[string]string) error
	HandleAppleCallback(ctx context.Context, notification *request.AppleIAPNotification) error
	VerifyAlipayNotifySign(ctx context.Context, bm gopay.BodyMap) error
	HandleStripeCallback(ctx context.Context, payload []byte, header string) error
}

type logic struct {
	repo  repo.Factory
	cache cache.Factory
	srv   service.Factory
}

func NewCallback(repo repo.Factory, cache cache.Factory, service service.Factory) Logic {
	return &logic{
		repo:  repo,
		cache: cache,
		srv:   service,
	}
}

package logic

import (
	"microservices/cache"
	"microservices/logic/auth"
	"microservices/logic/callback"
	"microservices/logic/file"
	"microservices/logic/goods"
	"microservices/logic/notify"
	"microservices/logic/order"
	"microservices/logic/user"
	"microservices/repo"
	"microservices/service"
)

// Factory defines functions used to return resource interface.
type Factory interface {
	User() user.Logic
	Auth() auth.Logic
	Callback() callback.Logic
	Notify() notify.Logic
	Order() order.Logic
	Goods() goods.Logic
	File() file.Logic
}

type factory struct {
	repo    repo.Factory
	cache   cache.Factory
	service service.Factory
}

func (l *factory) Notify() notify.Logic {
	return notify.NewNotify(l.repo, l.cache, l.service)
}

func (l *factory) User() user.Logic {
	return user.NewLogic(l.repo, l.cache, l.service)
}

func (l *factory) Auth() auth.Logic {
	return auth.NewAuth(l.repo, l.cache, l.service)
}

func (l *factory) Callback() callback.Logic {
	return callback.NewCallback(l.repo, l.cache, l.service)
}

func (l *factory) Order() order.Logic {
	return order.NewLogic(l.repo, l.cache, l.service)
}

func (l *factory) Goods() goods.Logic {
	return goods.NewLogic(l.repo, l.cache, l.service)
}

func (l *factory) File() file.Logic {
	return file.NewLogic(l.repo, l.cache, l.service)
}

// NewLogic .
func NewLogic(repo repo.Factory, cache cache.Factory, service service.Factory) Factory {
	return &factory{repo, cache, service}
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/xxx-newbee/go-micro/gateway/internal/adapter"
	"github.com/xxx-newbee/go-micro/gateway/internal/config"
	"github.com/xxx-newbee/go-micro/gateway/internal/types"
	"github.com/zeromicro/go-zero/core/breaker"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserService types.UserService
	UserBreaker breaker.Breaker
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.UserRpcConf)
	clientService := adapter.NewUserClientAdapter(client)
	bk := breaker.NewBreaker(breaker.WithName("gateway"))

	return &ServiceContext{
		Config:      c,
		UserService: clientService,
		UserBreaker: bk,
	}
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/xxx-newbee/gateway/internal/adapter"
	"github.com/xxx-newbee/gateway/internal/config"
	"github.com/xxx-newbee/gateway/internal/middleware"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserService  types.UserService
	UserBreaker  breaker.Breaker
	JwtAuth      rest.Middleware
	RateLimiter  rest.Middleware
	RequestTimer rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := zrpc.MustNewClient(c.UserRpcConf)
	clientService := adapter.NewUserClientAdapter(client)
	bk := breaker.NewBreaker(breaker.WithName("gateway"))

	return &ServiceContext{
		Config:       c,
		UserService:  clientService,
		UserBreaker:  bk,
		JwtAuth:      middleware.NewJwtAuthMiddleware(c.JWT.Secret, c.JWT.AccessExpire, c.JWT.RefreshExpire).Handle,
		RateLimiter:  middleware.NewRateLimiterMiddleware().Handle,
		RequestTimer: middleware.NewRequestTimerMiddleware().Handle,
	}
}

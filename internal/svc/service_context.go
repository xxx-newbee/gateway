// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/xxx-newbee/chat/chat"
	"github.com/xxx-newbee/gateway/internal/config"
	"github.com/xxx-newbee/gateway/internal/middleware"
	"github.com/xxx-newbee/order/order"
	"github.com/xxx-newbee/user/user"
	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      user.UserClient
	OrderRpc     order.OrderClient
	ChatRpc      chat.ChatClient
	UserBreaker  breaker.Breaker
	JwtAuth      rest.Middleware
	RateLimiter  rest.Middleware
	RequestTimer rest.Middleware
	Header       rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	user_client := zrpc.MustNewClient(c.UserRpcConf)
	order_client := zrpc.MustNewClient(c.OrderRpcConf)
	chat_client := zrpc.MustNewClient(c.ChatRpcConf)
	bk := breaker.NewBreaker(breaker.WithName("gateway"))

	return &ServiceContext{
		Config:       c,
		UserRpc:      user.NewUserClient(user_client.Conn()),
		OrderRpc:     order.NewOrderClient(order_client.Conn()),
		ChatRpc:      chat.NewChatClient(chat_client.Conn()),
		UserBreaker:  bk,
		JwtAuth:      middleware.NewJwtAuthMiddleware(c.JWT.Secret, c.JWT.AccessExpire, c.JWT.RefreshExpire).Handle,
		RateLimiter:  middleware.NewRateLimiterMiddleware().Handle,
		RequestTimer: middleware.NewRequestTimerMiddleware().Handle,
		Header:       middleware.NewHeaderMiddleware().Handle,
	}
}

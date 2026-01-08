// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf zrpc.RpcClientConf
	Routes []Route
	JWT struct {
		Secret string
		AccessExpire int64
		RefreshExpire int64
	}
}

type Route struct {
	Name string
	Path string
	Method string
	Dest string
}

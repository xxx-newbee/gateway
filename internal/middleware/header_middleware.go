// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

type HeaderMiddleware struct {
}

func NewHeaderMiddleware() *HeaderMiddleware {
	return &HeaderMiddleware{}
}

func (m *HeaderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 请求头传递给下游服务
		md := metadata.MD{}
		// metadata不能常规命名HTTP协议的请求头，在rpc协议中会被替换，比如：User-Agent
		md["UA"] = []string{r.Header.Get("User-Agent")}
		md["remote-addr"] = []string{r.RemoteAddr}
		md["Authorization"] = []string{r.Header.Get("Authorization")}
		ctx := metadata.NewOutgoingContext(r.Context(), md)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

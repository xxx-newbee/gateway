// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RequestTimerMiddleware struct {
}

func NewRequestTimerMiddleware() *RequestTimerMiddleware {
	return &RequestTimerMiddleware{}
}

func (m *RequestTimerMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		elapsed := time.Since(start)
		logx.Infof("请求：%s %s 处理时间：%v", r.Method, r.URL.Path, elapsed)
	}
}

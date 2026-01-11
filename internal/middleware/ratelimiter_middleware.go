// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RateLimiterMiddleware struct {
	requestRecords map[string][]time.Time
	mutex          sync.RWMutex
	maxRequests    int
	windowSize     time.Duration
}

func NewRateLimiterMiddleware() *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		requestRecords: make(map[string][]time.Time),
		maxRequests:    1,
		windowSize:     time.Second * 10,
	}
}

func (m *RateLimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		clientIP := r.RemoteAddr
		m.mutex.Lock()
		defer m.mutex.Unlock()

		now := time.Now()

		// 清理过期的请求记录
		if records, exists := m.requestRecords[clientIP]; exists {
			var validRecords []time.Time
			for _, t := range records {
				if now.Sub(t) <= m.windowSize {
					validRecords = append(validRecords, t)
				}
			}
			m.requestRecords[clientIP] = validRecords
		}
		// 检查请求频率
		if len(m.requestRecords[clientIP]) >= m.maxRequests {
			httpx.Error(w, errors.New("请求过于频繁，请稍后再试"))
		}
		// 记录本次请求
		m.requestRecords[clientIP] = append(m.requestRecords[clientIP], now)

		next(w, r)
	}
}

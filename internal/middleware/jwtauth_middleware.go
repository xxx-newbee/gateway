// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtAuthMiddleware struct {
	secret        string
	accessExpire  int64
	refreshExpire int64
}

func NewJwtAuthMiddleware(screct string, accessExpire int64, refreshExpire int64) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		secret:        screct,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokenStr string
		// 获取浏览器默认cookie
		cookie, err := r.Cookie("access_token")
		if err == nil {
			tokenStr = cookie.Value
		}

		if tokenStr == "" {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
				tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if tokenStr == "" {
			httpx.Error(w, errors.New("未登录，请先登录"))
			return
		}
		// 解析 token
		token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(m.secret), nil
		})

		if err != nil || !token.Valid {
			httpx.Error(w, errors.New("无效的token，请重新登录"))
			return
		}
		_, ok := token.Claims.(*JWTClaims)
		if !ok {
			httpx.Error(w, errors.New("无法适配token，请重新登录"))
			return
		}

		next(w, r)
	}
}

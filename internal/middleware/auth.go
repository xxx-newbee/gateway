package middleware

import (
	"context"
	"errors"
	"github.com/xxx-newbee/go-micro/gateway/internal/svc"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JWTClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AuthMiddleware(srv *svc.ServiceContext) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var tokenStr string
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

			token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(srv.Config.JWT.Secret), nil
			})

			if err != nil || !token.Valid {
				httpx.Error(w, errors.New("无效的token，请重新登录"))
				return
			}
			claims, ok := token.Claims.(*JWTClaims)
			if !ok {
				httpx.Error(w, errors.New("无法适配token，请重新登录"))
				return
			}
			ctx := context.WithValue(r.Context(), "userID", claims.UserID)
			ctx = context.WithValue(ctx, "username", claims.Username)
			ctx = context.WithValue(ctx, "Authorization", tokenStr)
			r = r.WithContext(ctx)

			next(w, r)
		}
	}
}

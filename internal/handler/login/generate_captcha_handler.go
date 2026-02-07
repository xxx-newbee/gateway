// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"net/http"

	"github.com/xxx-newbee/gateway/internal/logic/login"
	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := login.NewGenerateCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GenerateCaptcha()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

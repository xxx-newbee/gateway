// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"net/http"

	"github.com/xxx-newbee/gateway/internal/logic/user"
	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteLoginLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteLoginLogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteLoginLogLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLoginLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

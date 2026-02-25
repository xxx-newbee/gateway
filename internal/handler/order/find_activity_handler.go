// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order

import (
	"net/http"

	"github.com/xxx-newbee/gateway/internal/logic/order"
	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FindActivityHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := order.NewFindActivityLogic(r.Context(), svcCtx)
		resp, err := l.FindActivity()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

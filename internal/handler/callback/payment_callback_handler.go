package callback

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xxx-newbee/gateway/internal/logic/callback"
	"github.com/xxx-newbee/gateway/internal/svc"
)

func PaymentCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"code": "FAIL", "message": "read body failed"})
			return
		}
		defer r.Body.Close()

		l := callback.NewPaymentCallbackLogic(r.Context(), svcCtx)
		resp, err := l.PaymentCallback(&callback.PaymentCallbackReq{
			Signature: r.Header.Get("Wechatpay-Signature"),
			Nonce:     r.Header.Get("Wechatpay-Nonce"),
			Timestamp: r.Header.Get("Wechatpay-Timestamp"),
			Serial:    r.Header.Get("Wechatpay-Serial"),
			Body:      string(bodyBytes),
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"code": "FAIL", "message": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

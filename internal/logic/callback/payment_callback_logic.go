package callback

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentCallbackLogic {
	return &PaymentCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type PaymentCallbackReq struct {
	Signature string
	Nonce     string
	Timestamp string
	Serial    string
	Body      string
}

func (l *PaymentCallbackLogic) PaymentCallback(req *PaymentCallbackReq) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.PaymentCallbackResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.ProcessPaymentCallback(l.ctx, &orderPb.PaymentCallbackRequest{
			WechatpaySignature: req.Signature,
			WechatpayNonce:     req.Nonce,
			WechatpayTimestamp: req.Timestamp,
			WechatpaySerial:    req.Serial,
			Body:               req.Body,
		})
		return innerErr
	}, func(err error) bool {
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		return &types.BaseResponse{Code: 500, Msg: err.Error()}, nil
	}
	return &types.BaseResponse{Code: 200, Msg: "ok", Data: rpcResp}, nil
}

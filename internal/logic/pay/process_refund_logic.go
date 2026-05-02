package pay

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessRefundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessRefundLogic {
	return &ProcessRefundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessRefundLogic) ProcessRefund(req *types.RefundPayRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.RefundResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.Refund(l.ctx, &orderPb.RefundRequest{
			OrderNo:      req.OrderNo,
			RefundAmount: int64(req.RefundAmount * 100),
			RefundReason: req.RefundReason,
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

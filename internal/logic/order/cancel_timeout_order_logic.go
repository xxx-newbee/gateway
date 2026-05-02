package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelTimeoutOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelTimeoutOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelTimeoutOrderLogic {
	return &CancelTimeoutOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelTimeoutOrderLogic) CancelTimeoutOrder(req *types.CancelTimeoutOrderRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.CancelTimeoutOrderResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.CancelTimeoutOrder(l.ctx, &orderPb.CancelTimeoutOrderRequest{
			OrderNo: req.OrderNo,
		})
		return innerErr
	}, func(err error) bool {
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		return &types.BaseResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.BaseResponse{
		Code: 200,
		Msg:  "ok",
		Data: rpcResp,
	}, nil
}

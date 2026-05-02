package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillOrderLogic {
	return &GetSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSeckillOrderLogic) GetSeckillOrder(req *types.GetSeckillOrderRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.GetSeckillOrderResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.GetSeckillOrder(l.ctx, &orderPb.GetSeckillOrderRequest{
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

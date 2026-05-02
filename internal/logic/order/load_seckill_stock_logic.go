package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoadSeckillStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoadSeckillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoadSeckillStockLogic {
	return &LoadSeckillStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoadSeckillStockLogic) LoadSeckillStock(req *types.LoadSeckillStockRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.LoadSeckillStockResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.LoadSeckillStock(l.ctx, &orderPb.LoadSeckillStockRequest{
			ActivityId: uint32(req.ActivityId),
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

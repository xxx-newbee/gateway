package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSeckillActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillActivityLogic {
	return &CreateSeckillActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillActivityLogic) CreateSeckillActivity(req *types.CreateSeckillActivityRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.CreateSeckillActivityResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.CreateSeckillActivity(l.ctx, &orderPb.CreateSeckillActivityRequest{
			ProductId:    req.ProductId,
			SeckillPrice: req.SeckillPrice,
			StockNum:     req.StockNum,
			StartTime:    req.StartTime,
			EndTime:      req.EndTime,
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

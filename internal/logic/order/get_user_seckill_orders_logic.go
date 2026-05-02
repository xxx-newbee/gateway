package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSeckillOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserSeckillOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSeckillOrdersLogic {
	return &GetUserSeckillOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSeckillOrdersLogic) GetUserSeckillOrders(req *types.GetUserSeckillOrdersRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.GetUserSeckillOrdersResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.GetUserSeckillOrders(l.ctx, &orderPb.GetUserSeckillOrdersRequest{
			UserId:   req.UserId,
			Page:     req.Page,
			PageSize: req.PageSize,
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

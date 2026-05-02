package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	orderPb "github.com/xxx-newbee/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindActivityLogic {
	return &FindActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindActivityLogic) FindActivity(req *types.FindActivityRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *orderPb.GetSeckillActivityResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.OrderRpc.GetSeckillActivity(l.ctx, &orderPb.GetSeckillActivityRequest{
			ActivityId: uint32(req.Id),
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
		Data: map[string]interface{}{
			"activity_id":   rpcResp.ActivityId,
			"product_id":    rpcResp.ProductId,
			"seckill_price": rpcResp.SeckillPrice,
			"stock_num":     rpcResp.StockNum,
			"surplus_stock": rpcResp.SurplusStock,
			"start_time":    rpcResp.StartTime,
			"end_time":      rpcResp.EndTime,
			"status":        rpcResp.Status,
			"success":       rpcResp.Success,
			"msg":           rpcResp.Msg,
		},
	}, nil
}

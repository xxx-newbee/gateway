package product

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	productPb "github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSkuLogic {
	return &UpdateSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSkuLogic) UpdateSku(req *types.UpdateSkuRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *productPb.UpdateSkuResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.ProductRpc.UpdateSku(l.ctx, &productPb.UpdateSkuRequest{
			SkuId:   req.SkuId,
			SkuCode: req.SkuCode,
			Specs:   req.Specs,
			Price:   req.Price,
			Stock:   req.Stock,
			Status:  req.Status,
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

package product

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	productPb "github.com/xxx-newbee/product/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSkuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSkuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSkuLogic {
	return &DeleteSkuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSkuLogic) DeleteSku(req *types.DeleteSkuRequest) (resp *types.BaseResponse, err error) {
	var rpcResp *productPb.DeleteSkuResponse
	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		rpcResp, innerErr = l.svcCtx.ProductRpc.DeleteSku(l.ctx, &productPb.DeleteSkuRequest{
			SkuId: req.SkuId,
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

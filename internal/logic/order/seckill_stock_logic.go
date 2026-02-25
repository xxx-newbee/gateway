// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillStockLogic {
	return &SeckillStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillStockLogic) SeckillStock(req *types.SeckillRequest) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

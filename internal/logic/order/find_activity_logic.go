// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package order

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"

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

func (l *FindActivityLogic) FindActivity() (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegistRequest) (*types.BaseResponse, error) {
	var resp *types.RegistResponse
	err := l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		resp, innerErr = l.svcCtx.UserService.Register(l.ctx, req)
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
		Data: resp,
	}, nil

}

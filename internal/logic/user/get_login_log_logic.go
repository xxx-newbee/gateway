// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/xxx-newbee/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginLogLogic {
	return &GetLoginLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginLogLogic) GetLoginLog(req *types.LoginLogRequest) (*types.BaseResponse, error) {
	// todo: add your logic here and delete this line
	var res *user.LoginLogResponse
	err := l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		res, innerErr = l.svcCtx.UserRpc.GetLoginLog(l.ctx, &user.GetLoginLogRequest{Page: req.Page})
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
		Data: res.LoginLogs,
		Msg:  "success",
	}, nil
}

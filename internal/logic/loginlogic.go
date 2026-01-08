// gateway/internal/logic/loginlogic.go
package logic

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	var resp *types.LoginResponse

	// 熔断器保护， 失败时降级
	err := l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		resp, innerErr = l.svcCtx.UserService.Login(l.ctx, req)
		return innerErr
	}, func(err error) bool {
		// 触发熔断
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		// 降级处理
		return &types.LoginResponse{
			Token:    "",
			UserId:   0,
			Username: "匿名用户",
			Nickname: "匿名",
		}, nil
	}

	return resp, nil
}

func (l *LoginLogic) Register(req *types.RegistRequest) (*types.RegistResponse, error) {
	resp, err := l.svcCtx.UserService.Regist(l.ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

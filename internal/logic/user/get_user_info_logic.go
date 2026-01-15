// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"errors"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"google.golang.org/grpc/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (*types.GetUserInfoResponse, error) {

	tokenStr, ok := l.ctx.Value("Authorization").(string)
	if !ok || tokenStr == "" {
		return nil, errors.New("authorization token not found in context")
	}

	md := metadata.Pairs("Authorization", tokenStr)
	l.ctx = metadata.NewOutgoingContext(l.ctx, md)

	var resp *types.GetUserInfoResponse
	err := l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		resp, innerErr = l.svcCtx.UserService.GetUserInfo(l.ctx)
		return innerErr
	}, func(err error) bool {
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

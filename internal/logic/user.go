package logic

import (
	"context"
	"errors"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/metadata"
)

type User struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUser(svc *svc.ServiceContext, ctx context.Context) *User {
	return &User{
		ctx:    ctx,
		svcCtx: svc,
		Logger: logx.WithContext(ctx),
	}
}

func (u *User) GetUserInfo(req *types.GetUserRequest) (*types.GetUserResponse, error) {
	tokenStr, ok := u.ctx.Value("Authorization").(string)
	if !ok || tokenStr == "" {
		return nil, errors.New("authorization token not found in context")
	}

	md := metadata.Pairs("Authorization", tokenStr)
	u.ctx = metadata.NewOutgoingContext(u.ctx, md)

	var resp *types.GetUserResponse
	err := u.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		resp, innerErr = u.svcCtx.UserService.GetUserInfo(u.ctx, req)
		return innerErr
	}, func(err error) bool {
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		return &types.GetUserResponse{
			Username: "匿名用户",
			Nickname: "匿名",
		}, nil
	}

	return resp, nil
}

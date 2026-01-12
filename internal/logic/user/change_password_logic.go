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

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.UpdateResponse, err error) {
	// todo: add your logic here and delete this line
	tokenStr, ok := l.ctx.Value("Authorization").(string)
	if !ok || tokenStr == "" {
		return nil, errors.New("authorization token not found in context")
	}

	md := metadata.Pairs("Authorization", tokenStr)
	l.ctx = metadata.NewOutgoingContext(l.ctx, md)

	err = l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
		var innerErr error
		resp, innerErr = l.svcCtx.UserService.ChangePassword(l.ctx, req)
		return innerErr
	}, func(err error) bool {
		return err != nil && context.DeadlineExceeded == err
	})

	if err != nil {
		return &types.UpdateResponse{Status: "fail", Msg: err.Error()}, nil
	}

	return
}

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

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (*types.BaseResponse, error) {
	// todo: add your logic here and delete this line
	tokenStr, ok := l.ctx.Value("Authorization").(string)
	if !ok || tokenStr == "" {
		return nil, errors.New("authorization token not found in context")
	}

	md := metadata.Pairs("Authorization", tokenStr)
	l.ctx = metadata.NewOutgoingContext(l.ctx, md)

	err := l.svcCtx.UserService.ChangePassword(l.ctx, req)

	if err != nil {
		return &types.BaseResponse{
			Code: 500,
			Msg:  err.Error(),
		}, nil
	}

	return &types.BaseResponse{
		Code: 200,
		Msg:  "ok",
	}, nil
}

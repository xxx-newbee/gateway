// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"context"
	"strings"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/xxx-newbee/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendEmailRequest) (*types.BaseResponse, error) {
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		return &types.BaseResponse{Code: 200, Msg: "illegal email"}, nil
	}
	resp, err := l.svcCtx.UserRpc.SendEmail(l.ctx, &user.SendEmailRequest{Email: req.Email})

	if err != nil {
		return &types.BaseResponse{Code: 500, Msg: err.Error()}, nil
	}

	return &types.BaseResponse{Code: 200, Msg: "ok", Data: resp}, nil
}

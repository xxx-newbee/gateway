// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package login

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/svc"
	"github.com/xxx-newbee/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCaptchaLogic {
	return &GenerateCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateCaptchaLogic) GenerateCaptcha() (*types.BaseResponse, error) {
	//var resp *types.CaptchaResponse
	//err := l.svcCtx.UserBreaker.DoWithAcceptable(func() error {
	//	var innerErr error
	//	resp, innerErr = l.svcCtx.UserService.GenerateCaptcha(l.ctx)
	//	return innerErr
	//}, func(err error) bool {
	//	return err != nil && context.DeadlineExceeded == err
	//})
	resp, err := l.svcCtx.UserService.GenerateCaptcha(l.ctx)

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

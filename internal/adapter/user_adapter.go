package adapter

import (
	"context"

	"github.com/xxx-newbee/gateway/internal/types"
	"github.com/xxx-newbee/user/user"
	"github.com/xxx-newbee/user/userclient"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type UserClientAdapter struct {
	cli userclient.User
}

func NewUserClientAdapter(rpcClient zrpc.Client) *UserClientAdapter {
	return &UserClientAdapter{
		cli: userclient.NewUser(rpcClient),
	}
}

func (e *UserClientAdapter) Regist(ctx context.Context, in *types.RegistRequest, opts ...grpc.CallOption) (*types.RegistResponse, error) {
	protoReq := &user.RegisterRequest{
		Username:     in.Username,
		Password:     in.Password,
		Nickname:     in.Nickname,
		WalletAddr:   in.WalletAddr,
		ReferralCode: in.ReferralCode,
	}

	protoResp, err := e.cli.Register(ctx, protoReq)
	if err != nil {
		return nil, err
	}

	return &types.RegistResponse{
		Username:       protoResp.Username,
		Nickname:       protoResp.Nickname,
		WalletAddr:     protoResp.WalletAddr,
		UserRefferCode: protoResp.UserReferralCode,
		ReferralCode:   protoResp.ReferralCode,
	}, nil
}

func (e *UserClientAdapter) Login(ctx context.Context, in *types.LoginRequest, opts ...grpc.CallOption) (*types.LoginResponse, error) {
	protoReq := &user.LoginRequest{
		Username: in.Username,
		Password: in.Password,
	}
	protoResp, err := e.cli.Login(ctx, protoReq)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		UserId:           protoResp.UserId,
		Username:         protoResp.Username,
		Nickname:         protoResp.Nickname,
		Token:            protoResp.Token,
		WalletAddr:       protoResp.WalletAddr,
		UserReferralCode: protoResp.UserReferralCode,
		ReferralCode:     protoResp.ReferralCode,
	}, nil
}

func (e *UserClientAdapter) GetUserInfo(ctx context.Context, in *types.GetUserRequest, opts ...grpc.CallOption) (*types.GetUserResponse, error) {
	protoReq := &user.GetUserInfoRequest{}
	protoResp, err := e.cli.GetUserInfo(ctx, protoReq)
	if err != nil {
		return nil, err
	}
	return &types.GetUserResponse{
		Username:       protoResp.Username,
		Nickname:       protoResp.Nickname,
		WalletAddr:     protoResp.WalletAddr,
		UserRefferCode: protoResp.UserReferralCode,
		ReferralCode:   protoResp.ReferralCode,
	}, nil
}

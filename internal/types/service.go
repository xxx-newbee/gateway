package types

import (
	"context"

	"google.golang.org/grpc"
)

type UserService interface {
	Register(context.Context, *RegistRequest, ...grpc.CallOption) (*RegistResponse, error)
	Login(context.Context, *LoginRequest, ...grpc.CallOption) (*LoginResponse, error)
	GetUserInfo(context.Context, ...grpc.CallOption) (*GetUserInfoResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoReqest, ...grpc.CallOption) (*UpdateResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest, ...grpc.CallOption) (*UpdateResponse, error)
}

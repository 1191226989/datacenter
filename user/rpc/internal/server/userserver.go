// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"datacenter/user/rpc/internal/logic"
	"datacenter/user/rpc/internal/svc"
	"datacenter/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 获取用户数据（用户id）
func (s *UserServer) GetUserInfoById(ctx context.Context, in *user.UserIdRequest) (*user.UserInfoResponse, error) {
	l := logic.NewGetUserInfoByIdLogic(ctx, s.svcCtx)
	return l.GetUserInfoById(in)
}

// 获取用户数据（用户名）
func (s *UserServer) GetUserInfoByUsername(ctx context.Context, in *user.UserNameRequest) (*user.UserInfoResponse, error) {
	l := logic.NewGetUserInfoByUsernameLogic(ctx, s.svcCtx)
	return l.GetUserInfoByUsername(in)
}

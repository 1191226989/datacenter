package logic

import (
	"context"

	"datacenter/user/rpc/internal/svc"
	"datacenter/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUsernameLogic {
	return &GetUserInfoByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户数据（用户名）
func (l *GetUserInfoByUsernameLogic) GetUserInfoByUsername(in *user.UserNameRequest) (*user.UserInfoResponse, error) {
	userModel, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}

	return &user.UserInfoResponse{
		Id:       userModel.Id,
		Username: userModel.Username,
		Password: userModel.Password,
		Mobile:   userModel.Mobile,
		Nickname: userModel.Nickname,
	}, nil
}

package public

import (
	"context"

	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestingLogic {
	return &TestingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestingLogic) Testing(req *types.TestingRequest) (resp *types.TestingResponse, err error) {
	return &types.TestingResponse{
		Message: "测试访问成功",
	}, nil
}

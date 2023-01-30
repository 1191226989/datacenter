package game

import (
	"context"

	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameDrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameDrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameDrawLogic {
	return &GameDrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameDrawLogic) GameDraw(req *types.GameDrawRequest) (resp *types.GameDrawResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

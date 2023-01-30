package game

import (
	"context"

	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameInfoLogic {
	return &GameInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameInfoLogic) GameInfo(req *types.GameInfoRequest) (resp *types.GameInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

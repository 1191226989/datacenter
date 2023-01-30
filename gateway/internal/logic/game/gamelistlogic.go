package game

import (
	"context"

	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GameListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameListLogic {
	return &GameListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameListLogic) GameList(req *types.GameListRequest) (resp *types.GameListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

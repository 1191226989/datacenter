package game

import (
	"net/http"

	"datacenter/gateway/internal/logic/game"
	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GameDrawHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GameDrawRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := game.NewGameDrawLogic(r.Context(), svcCtx)
		resp, err := l.GameDraw(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

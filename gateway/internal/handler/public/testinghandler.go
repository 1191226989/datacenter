package public

import (
	"net/http"

	"datacenter/gateway/internal/logic/public"
	"datacenter/gateway/internal/svc"
	"datacenter/gateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TestingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestingRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewTestingLogic(r.Context(), svcCtx)
		resp, err := l.Testing(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

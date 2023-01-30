package middleware

import (
	"datacenter/gateway/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type CheckPassMiddleware struct {
}

func NewCheckPassMiddleware() *CheckPassMiddleware {
	return &CheckPassMiddleware{}
}

func (m *CheckPassMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	logx.Info("public check pass middleware")
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("pass") != "123456" {
			httpx.OkJson(w, &types.TestingResponse{
				Message: "pass is required",
			})
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}

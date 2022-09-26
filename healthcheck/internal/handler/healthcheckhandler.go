package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"syndicate/healthcheck/internal/logic"
	"syndicate/healthcheck/internal/svc"
	"syndicate/healthcheck/internal/types"
)

func HealthCheckHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHealthCheckLogic(r.Context(), ctx)
		resp, err := l.HealthCheck(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

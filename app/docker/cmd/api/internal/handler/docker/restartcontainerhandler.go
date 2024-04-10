package docker

import (
	common_models "go-zero-container/common/global/models"
	"go-zero-container/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/docker/cmd/api/internal/logic/docker"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

func RestartContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req common_models.ContainerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := docker.NewRestartContainerLogic(r.Context(), svcCtx)
		err := l.RestartContainer(&req)
		result.HttpResult(r, w, nil, err)
	}
}

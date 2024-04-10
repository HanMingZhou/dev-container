package docker

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/docker/cmd/api/internal/logic/docker"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	common_models "go-zero-container/common/global/models"
	"go-zero-container/common/result"
	"net/http"
)

func StartContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req common_models.ContainerReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := docker.NewStartContainerLogic(r.Context(), svcCtx)
		err := l.StartContainer(&req)
		result.HttpResult(r, w, nil, err)
	}
}

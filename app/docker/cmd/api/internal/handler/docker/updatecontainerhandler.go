package docker

import (
	"go-zero-container/common/global/models"
	"go-zero-container/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/docker/cmd/api/internal/logic/docker"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

func UpdateContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.UpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := docker.NewUpdateContainerLogic(r.Context(), svcCtx)
		err := l.UpdateContainer(&req, r)
		result.HttpResult(r, w, nil, err)
	}
}

package docker

import (
	"go-zero-container/common/global/models"
	"go-zero-container/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/docker/cmd/api/internal/logic/docker"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

func RenameContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RenameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := docker.NewRenameContainerLogic(r.Context(), svcCtx)
		err := l.RenameContainer(&req)
		result.HttpResult(r, w, nil, err)

	}
}

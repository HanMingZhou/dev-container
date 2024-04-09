package image

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/image/cmd/api/internal/logic/image"
	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"
)

func GetUserImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := image.NewGetUserImageLogic(r.Context(), svcCtx)
		resp, err := l.GetUserImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

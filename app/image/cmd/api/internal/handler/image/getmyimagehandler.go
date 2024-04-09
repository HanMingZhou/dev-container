package image

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-container/app/image/cmd/api/internal/logic/image"
	"go-zero-container/app/image/cmd/api/internal/svc"
)

func GetMyImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := image.NewGetMyImageLogic(r.Context(), svcCtx)
		resp, err := l.GetMyImage()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

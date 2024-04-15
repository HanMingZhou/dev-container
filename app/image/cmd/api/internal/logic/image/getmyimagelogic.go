package image

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"
	"go-zero-container/app/image/cmd/rpc/imageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyImageLogic {
	return &GetMyImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyImageLogic) GetMyImage() (resp *types.GetMyImageResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.ImageRpc.GetMyImage(l.ctx, &imageserver.GetMyImageReq{})
	if err != nil {
		return nil, errors.New("getMyImage failed")
	}
	var imageList []types.ImageRegistry
	for _, img := range res.ImageRegistry {
		imageList = append(imageList, types.ImageRegistry{
			Rid:            img.Rid,
			Kind:           img.Kind,
			UserId:         cast.ToUint(img.UserId),
			Name:           img.Name,
			Url:            img.Url,
			Authentication: img.Authentication,
			Username:       img.Username,
			Password:       img.Password,
		})
	}
	return &types.GetMyImageResp{
		ImageList: imageList,
	}, nil

}

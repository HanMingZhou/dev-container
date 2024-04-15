package image

import (
	"context"
	"errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"
	"go-zero-container/app/image/cmd/rpc/imageserver"
)

type GetAllImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllImageLogic {
	return &GetAllImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllImageLogic) GetAllImage() (resp *types.GetAllImageResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.ImageRpc.GetMyImage(l.ctx, &imageserver.GetMyImageReq{})
	if err != nil {
		return nil, errors.New("GetAllImage error: " + err.Error())
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
	return &types.GetAllImageResp{
		ImageList: imageList,
	}, nil

}

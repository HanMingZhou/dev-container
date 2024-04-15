package image

import (
	"context"
	"github.com/spf13/cast"
	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"
	"go-zero-container/app/image/cmd/rpc/imageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserImageLogic {
	return &GetUserImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserImageLogic) GetUserImage(req *types.GetUserImageReq) (resp *types.GetUserImageResp, err error) {
	// todo: add your logic here and delete this line
	//var userId uint
	//re := l.svcCtx.DB.Where("user_id = ?", req.UserId).Find(&userId, req.UserId)
	// 暂时先不过滤user.ID
	//if re.RowsAffected == 0 {
	//	return nil, errors.New("用户不存在")
	//}
	res, err := l.svcCtx.ImageRpc.GetUserImage(l.ctx, &imageserver.GetUserImageReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
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
	return &types.GetUserImageResp{
		ImageList: imageList,
	}, nil

}

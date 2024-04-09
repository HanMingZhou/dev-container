package image

import (
	"context"
	"errors"
	common_models "go-zero-container/common/global/models"
	"gorm.io/gorm"

	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"

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
	var imageRegistry []common_models.ImageRegistry
	//var userId uint
	//re := l.svcCtx.DB.Where("user_id = ?", req.UserId).Find(&userId, req.UserId)
	// 暂时先不过滤user.ID
	//if re.RowsAffected == 0 {
	//	return nil, errors.New("用户不存在")
	//}
	err = l.svcCtx.DB.Where("user_id = ? or kind = 1", req.UserId).Find(&imageRegistry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("当前用户未创建镜像")
	} else if err != nil {
		return nil, errors.New("查询失败，请稍后重试")
	}
	return &types.GetUserImageResp{
		ImageList: imageRegistry,
	}, nil

}

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

	var imageRegistry []common_models.ImageRegistry

	err = l.svcCtx.DB.Find(&imageRegistry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("未查询到镜像")
	} else if err != nil {
		return nil, errors.New("查询失败，请稍后重试")
	}
	return &types.GetAllImageResp{
		ImageList: imageRegistry,
	}, nil
}

package image

import (
	"context"
	"errors"
	"go-zero-container/common/ctxdata"
	common_models "go-zero-container/common/global/models"
	"gorm.io/gorm"

	"go-zero-container/app/image/cmd/api/internal/svc"
	"go-zero-container/app/image/cmd/api/internal/types"

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
	var imageRegistry []common_models.ImageRegistry

	err = l.svcCtx.DB.Where("user_id = ?", ctxdata.GetUidFromCtx(l.ctx)).Find(&imageRegistry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("当前用户未创建镜像")
	} else if err != nil {
		return nil, errors.New("查询失败，请稍后重试")
	}
	return &types.GetMyImageResp{
		ImageList: imageRegistry,
	}, nil

}

package logic

import (
	"context"
	"errors"
	"go-zero-container/common/ctxdata"
	common_models "go-zero-container/common/global/models"
	"gorm.io/gorm"

	"go-zero-container/app/image/cmd/rpc/internal/svc"
	"go-zero-container/app/image/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMyImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyImageLogic {
	return &GetMyImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMyImageLogic) GetMyImage(in *pb.GetMyImageReq) (*pb.GetMyImageResp, error) {
	// todo: add your logic here and delete this line
	var imageRegistry []common_models.ImageRegistry

	err := l.svcCtx.DB.Where("user_id = ?", ctxdata.GetUidFromCtx(l.ctx)).Find(&imageRegistry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("当前用户未创建镜像")
	} else if err != nil {
		return nil, errors.New("查询失败，请稍后重试")
	}
	return &pb.GetMyImageResp{}, nil

}

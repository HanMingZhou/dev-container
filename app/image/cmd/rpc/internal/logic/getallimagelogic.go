package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"

	"go-zero-container/app/image/cmd/rpc/internal/svc"
	"go-zero-container/app/image/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllImageLogic {
	return &GetAllImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllImageLogic) GetAllImage(in *pb.GetAllImageReq) (*pb.GetAllImageResp, error) {
	// todo: add your logic here and delete this line
	var imageRegistry []*pb.ImageRegistry

	err := l.svcCtx.DB.Find(&imageRegistry).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("未查询到镜像")
	} else if err != nil {
		return nil, errors.New("查询失败，请稍后重试")
	}
	return &pb.GetAllImageResp{
		ImageRegistry: imageRegistry,
	}, nil
}

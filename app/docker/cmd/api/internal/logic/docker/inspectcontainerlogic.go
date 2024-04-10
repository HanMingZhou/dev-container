package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type InspectContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInspectContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InspectContainerLogic {
	return &InspectContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InspectContainerLogic) InspectContainer() error {
	// todo: add your logic here and delete this line

	return nil
}

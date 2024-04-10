package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type UpdateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContainerLogic {
	return &UpdateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContainerLogic) UpdateContainer() error {
	// todo: add your logic here and delete this line

	return nil
}

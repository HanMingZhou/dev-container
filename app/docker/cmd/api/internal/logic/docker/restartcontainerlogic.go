package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type RestartContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestartContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartContainerLogic {
	return &RestartContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestartContainerLogic) RestartContainer() error {
	// todo: add your logic here and delete this line

	return nil
}

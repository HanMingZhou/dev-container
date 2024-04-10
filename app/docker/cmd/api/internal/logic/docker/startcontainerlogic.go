package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type StartContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartContainerLogic {
	return &StartContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartContainerLogic) StartContainer() error {
	// todo: add your logic here and delete this line

	return nil
}

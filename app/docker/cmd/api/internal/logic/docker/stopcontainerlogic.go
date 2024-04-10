package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type StopContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopContainerLogic {
	return &StopContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopContainerLogic) StopContainer() error {
	// todo: add your logic here and delete this line

	return nil
}

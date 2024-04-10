package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type DownloadContainerLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadContainerLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadContainerLogsLogic {
	return &DownloadContainerLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadContainerLogsLogic) DownloadContainerLogs() error {
	// todo: add your logic here and delete this line

	return nil
}

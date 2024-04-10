package docker

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type DeleteContainerByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContainerByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContainerByIdsLogic {
	return &DeleteContainerByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteContainerByIdsLogic) DeleteContainerByIds() error {
	// todo: add your logic here and delete this line

	return nil
}

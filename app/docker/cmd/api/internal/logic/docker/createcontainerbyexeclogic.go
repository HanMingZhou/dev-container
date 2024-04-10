package docker

import (
	"context"

	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/app/docker/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContainerByExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContainerByExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContainerByExecLogic {
	return &CreateContainerByExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContainerByExecLogic) CreateContainerByExec() (resp *types.CreateContainerResp, err error) {
	// todo: add your logic here and delete this line

	return
}

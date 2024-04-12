package docker

import (
	"context"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

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

func (l *RestartContainerLogic) RestartContainer(req *models.ContainerReq) error {
	// todo: add your logic here and delete this line
	// 0 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer 初始化失败", zap.Error(err))
		return err
	}
	// 1 遍历container.ids
	for _, id := range req.Ids {
		// 重启容器
		// 192.168.0.53 endpointId = 2
		err = client.RestartContainer(req.EndpointId, id)
		if err != nil {
			logx.Error("重启容器失败", zap.Error(err), zap.String("container.id", id))
			return err
		}
		logx.Info(" 重启container成功")
		return err
	}
	return err
}

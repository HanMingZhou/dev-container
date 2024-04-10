package docker

import (
	"context"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

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

func (l *StartContainerLogic) StartContainer(req *models.ContainerReq) error {
	// todo: add your logic here and delete this line
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer初始化失败", zap.Error(err))
		return err
	}
	// 遍历container.ids
	for _, id := range req.Ids {
		//// 获取container信息
		//containerInfo, err := client.InspectContainer(id)
		//if err != nil {
		//	log.Error("获取容器信息失败", zap.Error(err))
		//	return err
		//}
		//log.Info("获取容器信息成功", zap.String("containerId", containerInfo.ID))

		// 启动容器
		// 目前默认endpointId = 2
		err = client.StartContainer(req.EndpointId, id)
		if err != nil {
			logx.Error("启动容器失败", zap.Error(err))
			return err
		}
		logx.Info("启动容器成功", zap.String("containerId", id))
	}
	return nil
}

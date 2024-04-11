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
		// start a new container by 节点nodeID, 容器containerId
		err = client.StartContainer(req.EndpointId, id)
		if err != nil {
			logx.Error("启动Container失败", zap.Error(err))
			return err
		}
		logx.Info("启动Container成功", zap.String("Id", id))
		// 端口重新转发
		//err = ConServiceV2.ReForwardPort(req.EndpointId, id, userName)
		//if err != nil {
		//	logx.Error("端口检查失败", zap.Error(err), zap.String("Id", id))
		//	return errors.New("容器启动成功,端口检查失败:" + err.Error())
		//}
		//logx.Info("端口检查成功", zap.String("Id", id))
	}

	return nil
}

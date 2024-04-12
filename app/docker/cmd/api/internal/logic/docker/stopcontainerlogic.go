package docker

import (
	"context"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

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

func (l *StopContainerLogic) StopContainer(req *models.ContainerReq) error {
	// todo: add your logic here and delete this line
	// 0 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer 初始化失败")
		return err
	}
	// 1 遍历container.ids
	for _, id := range req.Ids {

		// 端口回收
		//innerErr := ConServiceV2.PortRecovery(id)
		//if innerErr != nil && innerErr.Error() != "端口回收-容器无可回收端口" {
		//	logx.Error("端口回收失败", zap.Error(innerErr), zap.String("ContainerId", id))
		//	return innerErr
		//}
		//logx.Info("端口操作完成", zap.String("容器id", id))
		err = client.StopContainer(req.EndpointId, id)
		if err != nil {
			logx.Error("停止Container失败", zap.Error(err), zap.String("ContainerId", id))
			return err
		}
		logx.Info("停止Container成功", zap.String("容器id", id))
	}
	return nil

}

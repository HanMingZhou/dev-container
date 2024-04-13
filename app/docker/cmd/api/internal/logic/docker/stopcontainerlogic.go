package docker

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/common/global/models"
	"go.uber.org/zap"
	"time"
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
	//client, err := container.NewContainer()
	//if err != nil {
	//	logx.Error("Portainer 初始化失败")
	//	return err
	//}
	client := l.svcCtx.Portiner
	time.Sleep(1 * time.Second)
	for _, id := range req.Ids {
		// stop a new container by 节点nodeID, 容器containerId
		err := client.StopContainer(req.EndpointId, id)
		if err != nil {
			logx.Error("停止Container失败", zap.Error(err))
			return err
		}
		logx.Info("停止Container成功", zap.String("Id", id))
	}
	return nil

}

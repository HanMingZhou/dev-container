package docker

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/common/global/models"
	"go.uber.org/zap"
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
	//client, err := container.NewContainer()
	//if err != nil {
	//	logx.Error("Portainer 初始化失败", zap.Error(err))
	//	return err
	//}
	client := l.svcCtx.Portiner
	// 1 遍历container.ids
	var err error
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

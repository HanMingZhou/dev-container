package docker

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/common/global/models"
	"go.uber.org/zap"
)

type GetContainerLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContainerLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContainerLogsLogic {
	return &GetContainerLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContainerLogsLogic) GetContainerLogs(req *models.ContainerLogReq) (logs []string, err error) {
	// todo: add your logic here and delete this line
	// 0 初始化portainer
	//client, err := container.NewContainer()
	//if err != nil {
	//	logx.Error("Portainer 初始化失败", zap.Error(err))
	//	return logs, err
	//}
	client := l.svcCtx.Portainer

	// 1 查看容器是否存在
	if err := l.svcCtx.DB.Where("container_id =?", req.Id).First(&models.Container{}).Error; err != nil {
		logx.Error("检查容器-查找容器失败", zap.Error(err))
		return nil, err
	}

	// 2 获取容器日志
	node := cast.ToInt32(req.Node)
	logDate, err := client.GetLogs(node, req.Id, req)
	if err != nil {
		logx.Error("获取容器日志失败", zap.Error(err))
		return nil, err
	}
	return logDate, nil
}

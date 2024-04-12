package docker

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"
	"strings"
)

type InspectContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInspectContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InspectContainerLogic {
	return &InspectContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InspectContainerLogic) InspectContainer(req *models.InspContainerReq) (resp *container.Inspect, err error) {
	// todo: add your logic here and delete this line
	nodeId := cast.ToInt32(req.Node)
	// 0 查看容器是否存在 查询Mysql
	if err := l.svcCtx.DB.Where("container_id = ?", req.ContainerId).First(&models.Container{}).Error; err != nil {
		logx.Error("检查容器-查找容器失败", zap.Error(err))
		// 目前连接本地数据库,存在查找容器id失败的情况,暂不处理
		return nil, err
	}

	// 1 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return nil, err
	}

	// 2 调用InspectContainer
	resp, err = client.InspectContainer(nodeId, req.ContainerId)
	if err != nil {
		logx.Error("Portainer检查Container失败", zap.Error(err))
		return
	}
	resp.Config.Env = nil
	result := strings.Split(resp.Config.Image, "/")
	resp.Config.Image = result[len(result)-1]
	return
}

package docker

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type DownloadContainerLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadContainerLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadContainerLogsLogic {
	return &DownloadContainerLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadContainerLogsLogic) DownloadContainerLogs(req *models.ContainerLogReq) error {
	// todo: add your logic here and delete this line

	// 0 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return err
	}

	// 1 查看容器是否存  table: container
	if err := l.svcCtx.DB.Where("container_id =?", req.Id).First(&models.Container{}).Error; err != nil {
		logx.Error("检查容器-查找容器失败", zap.Error(err))
		return err
	}
	// 2 根据l.ctx 获取username
	username := fmt.Sprintf("%v", l.ctx.Value("Username"))

	// 3 查看容器log
	nodeID := cast.ToInt32(req.Node)
	logData, err := client.GetLogs(nodeID, req.Id, req)
	if err != nil {
		logx.Error("获取容器日志失败", zap.Error(err))
		return err
	}

	// 4 保存日志
	// log文件的download路径根据yaml文件的LogPath修改
	logPath := fmt.Sprintf("%v/logs/%v/%v.log", l.svcCtx.Config.LogPath, username, req.Id)
	file, err := os.Create(logPath)
	if err != nil {
		logx.Error("创建日志文件失败", zap.String("logPath", logPath))
		return err
	}
	defer file.Close()
	for _, logInfo := range logData {
		_, err = file.WriteString(logInfo)
	}
	if err != nil {
		logx.Error("写入日志文件失败", zap.String("logPath", logPath))
		return err
	}
	logx.Info("写入日志成功", zap.String("logPath", logPath))

	return nil
}

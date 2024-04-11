package docker

import (
	"context"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type RenameContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRenameContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenameContainerLogic {
	return &RenameContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RenameContainerLogic) RenameContainer(req *models.RenameReq) error {
	// todo: add your logic here and delete this line
	// 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return err
	}
	// 获取username By l.ctx
	//userName := fmt.Sprintf("%s", l.ctx.Value("Username"))
	// 获取username By http.request
	args := make(map[string]string)
	conPrefix := l.svcCtx.Config.DockerAccount.ConPrefix
	args["name"] = l.svcCtx.Config.DockerAccount.ConPrefix + "-" + req.Name
	err = client.RenameContainer(req.EndpointId, req.ContainerId, args)
	if err != nil {
		logx.Error("更新容器名失败", zap.Error(err))
		return err
	}

	// 创建db
	db := l.svcCtx.DB.Model(&models.Container{})
	err = db.Where("container_id", req.ContainerId).Update("containe_name", conPrefix+"-"+req.Name).Error
	if err != nil {
		logx.Error(err.Error())
		return err
	}

	return nil
}

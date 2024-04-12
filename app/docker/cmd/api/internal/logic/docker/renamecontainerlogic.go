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
	// 0 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return err
	}
	// 1 获取username By l.ctx  目前更新container_name时暂时不用username
	// userName := fmt.Sprintf("%s", l.ctx.Value("Username"))
	// todo:暂时不通过http.request提供username
	// 获取username By http.request 也可以通过http.request时写入参数用来rename container

	// 2 根据DockerAccount.conprefix+请求的name 修改container name
	args := make(map[string]string)
	conPrefix := l.svcCtx.Config.DockerAccount.ConPrefix
	args["name"] = l.svcCtx.Config.DockerAccount.ConPrefix + "-" + req.Name
	err = client.RenameContainer(req.EndpointId, req.ContainerId, args)
	if err != nil {
		logx.Error("更新容器名失败", zap.Error(err))
		return err
	}

	// 4 更新数据库db: table “container”
	db := l.svcCtx.DB.Model(&models.Container{})
	err = db.Where("container_id", req.ContainerId).Update("containe_name", conPrefix+"-"+req.Name).Error
	if err != nil {
		logx.Error(err.Error())
		return err
	}

	return nil
}

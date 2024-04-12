package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/spf13/cast"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type UpdateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateContainerLogic {
	return &UpdateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateContainerLogic) UpdateContainer(req *models.UpdateReq, r *http.Request) error {
	// todo: add your logic here and delete this line

	// 0 初始化container.NewContainer()
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return err
	}
	// 1 解析restartPolicy
	logx.Error("UpdateContainerUpdateReq:", req)
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(&req)
	if err != nil {
		logx.Error("更新Container重启策略body解析失败", zap.Error(err))
		return err
	}

	// 2 获取节点ID,containerID by http.request
	nodeId := r.URL.Query().Get("nodeId")
	conId := r.URL.Query().Get("conId")

	// 3 更新restartPolicy
	err = client.UpdateContainer(cast.ToInt32(nodeId), conId, &buf)
	if err != nil {
		logx.Error("更新Container重启策略失败", zap.Error(err))
		return err
	}
	return nil
}

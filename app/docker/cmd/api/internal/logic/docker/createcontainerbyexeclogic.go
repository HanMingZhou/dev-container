package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type CreateContainerByExecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContainerByExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContainerByExecLogic {
	return &CreateContainerByExecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContainerByExecLogic) CreateContainerByExec(req *models.CreateExecReq) (execResp *models.CreateExecRsp, err error) {
	// todo: add your logic here and delete this line

	// 0 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer 初始化失败", zap.Error(err))
		return nil, err
	}
	// 1 查询containerID是否存在, table: container
	if err := l.svcCtx.DB.Where("container_id =?", req.ContainerId).First(&models.Container{}).Error; err != nil {
		logx.Error("检查容器-查找容器失败", zap.Error(err))
		return nil, err
	}
	// 2 定义CreateExecBody
	body := models.CreateExecBody{
		ID:           "",
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{req.Cmd},
	}

	// 3 Execute
	var buf bytes.Buffer
	logx.Error("CreateExecBody", body)
	// Encode writes the JSON encoding of v to the stream,
	// NewEncoder returns a new encoder that writes to w
	err = json.NewEncoder(&buf).Encode(&body)
	if err != nil {
		logx.Error("json 解析失败", zap.Error(err))
		return nil, err
	}
	logx.Error("CreateExecByExecBuf", buf)

	resp, err := client.ExecCreate(req.EndpointId, req.ContainerId, &buf)
	if err != nil {
		logx.Error("执行创建失败", zap.Error(err))
		return nil, err
	}
	logx.Error("执行创建成功", zap.String("containerId", req.ContainerId), zap.String("execId", resp.ExecId))

	return &models.CreateExecRsp{
		Token:  resp.Token,
		ExecId: resp.ExecId,
	}, nil
}

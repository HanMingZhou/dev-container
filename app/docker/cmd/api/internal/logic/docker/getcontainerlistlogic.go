package docker

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type GetContainerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContainerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContainerListLogic {
	return &GetContainerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetContainerListLogic) GetContainerList(req *models.ContainerSearch) (resp *models.GetContainerListResp, err error) {
	// todo: add your logic here and delete this line
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	// 创建db
	db := l.svcCtx.DB.Model(&models.Container{})
	var Cons []models.Container
	if req.StartCreatedAt != "" && req.EndCreatedAt != "" {
		layout := "2006-01-02 15:04:05"
		startCreatedAt, err := time.Parse(layout, req.StartCreatedAt)
		endCreatedAt, err := time.Parse(layout, req.EndCreatedAt)
		db = db.Where("created_at >=? and created_at <=?", startCreatedAt, endCreatedAt)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}

	if err != nil {
		logx.Error("获取Container总数失败", zap.Error(err))
		return
	}

	userUuid := fmt.Sprintf("%s", l.ctx.Value("UUID"))
	var total int64 = 0
	err = db.Where("user_uuid = ?", userUuid).Count(&total).Error
	fmt.Println("err=", err)
	logx.Error("userUuid:", userUuid)
	err = db.Where("user_uuid = ?", userUuid).Limit(limit).Offset(offset).Find(&Cons).Error
	if err != nil {
		logx.Error("获取Container列表失败", zap.Error(err))
		return
	}

	if len(Cons) == 0 {
		return nil, err
	}
	containersStatus, err := l.GetContainersStatusByUUID(Cons)
	if err != nil {
		logx.Error("获取Container-状态map失败", zap.Error(err))
		return
	}
	var list []models.Container
	for _, con := range Cons {
		result := strings.Split(con.Image, "/")
		//if result[len(result)-2] == l.svcCtx.Config.Harbor.Official {
		//	con.OfficialImage = true
		//}
		con.Image = result[len(result)-1]

		//con.Password = aes.AesDecrypt(con.Password, global.GVA_CONFIG.Aes.Key)
		con.Status = containersStatus[con.ContainerId]
		//con.PublicIp = global.GVA_CONFIG.Forward.PublicIP
		list = append(list, con)
	}

	return &models.GetContainerListResp{
		Total:         total,
		ContainerList: list,
		Page:          req.Page,
		PageSize:      req.PageSize,
	}, nil
}
func (l *GetContainerListLogic) GetContainersStatusByUUID(list []models.Container) (containersStatus map[string]string, err error) {
	// containersStatus[容器id] = 状态
	containersStatus = make(map[string]string, 0)
	// 1、构建map[节点id][]容器id
	nodeIds := make(map[string][]string, 0)
	for _, v := range list {
		nodeIds[v.NodeId] = append(nodeIds[v.NodeId], v.ContainerId)
	}

	// 2、查找所有节点的容器(调用portainer)
	// 2.1、初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return nil, err
	}
	for nodeId, cons := range nodeIds {
		containers, err := client.ListTargetContainers(nodeId, cons)
		if err != nil {
			logx.Error("获取containers失败", zap.Error(err))
			return nil, err
		}
		for _, c := range containers {
			containersStatus[c.ID] = c.State
		}
	}
	return containersStatus, nil

}

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

	// 0 设置偏移量
	limit := req.Pageinfo.Page
	offset := req.Pageinfo.PageSize
	logx.Error("page", limit)
	logx.Error("offset", offset)

	// 1 创建container db,查询数据库中符合条件的容器
	db := l.svcCtx.DB.Model(&models.Container{})
	if req.StartCreatedAt != "" && req.EndCreatedAt != "" {
		layout := "2006-01-02 15:04:05"
		startCreatedAt, err := time.Parse(layout, req.StartCreatedAt)
		endCreatedAt, err := time.Parse(layout, req.EndCreatedAt)
		db = db.Where("created_at >=? and created_at <=?", startCreatedAt, endCreatedAt)
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}
	//if err != nil {
	//	logx.Error("获取Container总数失败", zap.Error(err))
	//	return
	//}

	// 2 查询container数量by userUuid
	// uuid从l.ctx中获取,如果不用jwt token中获取uuid，则从数据库中暂时赋予一个临时值
	userUuid := fmt.Sprintf("%s", l.ctx.Value("UUID"))
	var total int64 = 0
	err = db.Where("user_uuid = ?", userUuid).Count(&total).Error
	logx.Error("userUuid:", userUuid)
	logx.Error("total:", total)

	// 3 从table container中获取所有的容器,并写入到cons中
	// 容器列表, type: slice
	var Cons []models.Container
	err = db.Where("user_uuid = ?", userUuid).Limit(limit).Offset(offset).Find(&Cons).Error
	if err != nil {
		logx.Error("获取Container列表失败", zap.Error(err))
		return
	}
	if len(Cons) == 0 {
		logx.Error("容器列表为:", len(Cons))
		return nil, err
	}

	// 4 查看容器的状态
	containersStatus, err := l.GetContainersStatusByUUID(Cons)
	if err != nil {
		logx.Error("获取Container-状态map失败", zap.Error(err))
		return
	}
	var list []models.Container
	// 5 遍历所有的cons,并从portainer中获取对应的state
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

	// 6 控制台输出
	fmt.Printf("container list: %v", list)

	return &models.GetContainerListResp{
		Total:         total,
		ContainerList: list,
		Page:          req.Pageinfo.Page,
		PageSize:      req.Pageinfo.PageSize,
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
	// 2.2、遍历所有节点下的所有容器
	for nodeId, cons := range nodeIds {
		containers, err := client.ListTargetContainers(nodeId, cons)
		if err != nil {
			logx.Error("获取containers失败", zap.Error(err))
			return nil, err
		}
		// 2.3、 遍历容器的state
		for _, c := range containers {
			containersStatus[c.ID] = c.State
		}
	}
	return containersStatus, nil

}

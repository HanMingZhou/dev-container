package docker

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-zero-container/common/global/models"
	"go-zero-container/common/utils/container"
	"go.uber.org/zap"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
)

type DeleteContainerByIdsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteContainerByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteContainerByIdsLogic {
	return &DeleteContainerByIdsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *DeleteContainerByIdsLogic) CheckContainerOwner(username string, containerID []string, nodeId int32) error {

	//根据username 在table container中查找所属的容器
	// 1 查找所有的容器
	var containers []models.Container
	if err := l.svcCtx.DB.Where("container_id in ?", containerID).Find(&containers).Error; err != nil {
		logx.Error("查找容器失败", zap.Error(err))
	}

	// 2 比较找到的容器和要删除出的容器
	logx.Error("删除的容器数量：len(containerID)=", len(containerID))
	logx.Error("查找到的容器数量：len(containers)=", len(containers))
	if len(containers) != len(containerID) {
		logx.Error("删除的容器数量与查找到的容器数量不一致")
		return fmt.Errorf("删除的容器数量与查找到的容器数量不一致")
	}

	// 3 检查节点
	//var node models.Node
	//if err := l.svcCtx.DB.Where("node_id =?", nodeId).First(&node).Error; err != nil {
	//	logx.Error("查找节点失败", zap.Error(err))
	//	return fmt.Errorf("查找节点失败")
	//}
	//logx.Error("查找的节点：", zap.Any("node", node))

	for _, con := range containers {
		if cast.ToInt32(con.NodeId) != nodeId {
			logx.Error("该容器不属于该节点")
			return fmt.Errorf("该容器不属于该节点")
		}
	}

	// 4 检查容器所有权
	var user models.SysUser
	logx.Error("username", "username")
	if err := l.svcCtx.DB.Where("username =?", username).First(&user).Error; err != nil {
		logx.Error("查找用户失败", zap.Error(err))
		return fmt.Errorf("查找用户失败")
	}
	logx.Error("查找的用户：", zap.Any("user", user))
	for _, con := range containers {
		if con.NickName != username {
			logx.Error("该容器不属于该用户")
			return fmt.Errorf("该容器不属于该用户,username:%v,contianer:%v", username, con.ContaineName)
		}
	}

	return nil
}

func (l *DeleteContainerByIdsLogic) DeleteContainerByIds(req *models.DeleteContainerReq) error {
	// todo: add your logic here and delete this line

	// 根据jwt上下文获取user_name
	username := fmt.Sprintf("%s", l.ctx.Value("Username"))

	// 初始化portainer
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return err
	}

	// 校验username所创建的container所有权 by username,nodeid, container_id
	err = l.CheckContainerOwner(username, req.Ids, req.EndpointId)
	if err != nil {
		logx.Error("删除容器校验失败", zap.Error(err))
		return err
	}
	logx.Error("删除容器校验成功")

	// 删除容器
	for _, id := range req.Ids {
		delErr := client.DeleteContainer(req.EndpointId, id)
		if delErr != nil {
			logx.Error("删除容器失败", zap.Error(delErr), zap.String("ContainerId", id))
			return delErr
		}
		logx.Info("删除容器成功", zap.String("ContainerId", id))
		// 删除container表中的记录
		delErr = l.svcCtx.DB.Delete(&[]models.Container{}, "container_id = ?", id).Error
		if delErr != nil {
			logx.Error("删除Container数据库操作失败", zap.Error(delErr), zap.String("container_id", id))
			return nil
		}
	}

	return nil
}

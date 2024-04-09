package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	ContainerId    string    `json:"containerId" form:"containerId" gorm:"column:container_id;comment:容器ID;"`
	ContaineName   string    `json:"containeName" form:"containeName" gorm:"column:containe_name;comment:容器名称;"`
	PortainerName  string    `json:"portainerName" form:"portainerName" gorm:"column:portainer_name;comment:portainer容器名称;"`
	Image          string    `json:"image" form:"image" gorm:"column:image;comment:容器镜像;"`
	NodeId         string    `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:GPU物理机_PortainerID;"`
	Status         string    `json:"status" form:"status" gorm:"column:status;comment:容器状态;"`
	NickName       string    `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:创建人名称;"`
	UserUuid       uuid.UUID `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:用户UUID;"`
	NodeIp         string    `json:"node_ip" form:"node_ip" gorm:"column:node_ip;comment:节点ip;"`
	PublicIp       string    `json:"public_ip" form:"public_ip" gorm:"column:public_ip;comment:容器的公网ip;"`
	Password       string    `json:"password" form:"password" gorm:"column:password;comment:容器密码;"`
	Gpus           string    `json:"gpus" form:"gpus" gorm:"type:text;column:gpus;comment:GPU占用数;"`
	DataVolumeName string    `json:"dataVolumeName" form:"dataVolumeName" gorm:"column:data_volume_name;comment:数据卷名称;"`
	OfficialImage  bool      `json:"officialImage" form:"officialImage" gorm:"-"`
	SharedStorage  bool      `json:"sharedStorage" form:"sharedStorage" gorm:"column:shared_storage;comment:是否有共享存储;"`
	SubscribeType  int       `json:"subscribeType" form:"subscribeType" gorm:"column:subscribe_type;comment:订阅类型(1、按时计费2、包时段制);"`
	Expired        bool      `json:"expired" form:"expired" gorm:"column:expired;comment:是否过期;"`
	CheckGpu       bool      `json:"checkGpu" form:"checkGpu" gorm:"column:check_gpu;comment:包时段制续费后,是否检测gpu;"`
	PortMap        string    `json:"portMap" form:"portMap" gorm:"column:port_map;comment:端口绑定;"`
	Option         string    `json:"option" form:"option" gorm:"column:option;comment:容器配置;"`
	ConvertStatus  int       `json:"convertStatus" form:"convertStatus" gorm:"column:convert_status;comment:转换状态;"`
}

//	默认关联查询表名设置为container

// 若不设置默认为 containers
func (Container) TableName() string {
	return "container"
}

package models

import "gorm.io/gorm"

// Node 节点结构体
type Node struct {
	gorm.Model
	// 节点名称
	Name string `json:"name" form:"name" gorm:"column:name;comment:节点的名称"`
	// IP地址
	IPAddress string `json:"ipAddress" form:"ip_address" gorm:"column:ip_address;comment:IP地址"`
	// 节点密码
	NodePassword string `json:"nodePassword" form:"node_password" gorm:"column:node_password;comment:节点密码"`
	// 节点PortainerID
	PortainerId string `json:"portainerId" form:"portainer_id" gorm:"column:portainer_id;comment:PortainerID"`
	// 节点状态
	NodeStatus string `json:"nodeStatus" form:"node_status" gorm:"column:node_status;comment:节点状态 0:未使用 1:使用中 2:异常"`
	// 节点模式
	NodeModel string `json:"nodeModel" form:"node_model" gorm:"column:node_model;comment:节点模式"`
	// GPU类型
	GpuModel string `json:"gpuModel" form:"gpu_model" gorm:"column:gpu_model;comment:GPU类型"`
	// 端口段编号
	NodePortNum uint `json:"nodePortNum" gorm:"column:node_port_num;comment:节点使用端口编码"`
}

// TableName 节点表名
func (Node) TableName() string {
	return "node"
}

// ShareNode 节点结构体
type ShareNode struct {
	gorm.Model
	// 节点名称
	Name string `json:"name" form:"name" gorm:"column:name;comment:节点的名称"`
	// IP地址
	IPAddress string `json:"ipAddress" form:"ip_address" gorm:"column:ip_address;comment:IP地址"`
	// 节点密码
	NodePassword string `json:"nodePassword" form:"node_password" gorm:"column:node_password;comment:节点密码"`
	// 节点PortainerID
	PortainerId string `json:"portainerId" form:"portainer_id" gorm:"column:portainer_id;comment:PortainerID"`
	// 节点状态
	NodeStatus string `json:"nodeStatus" form:"node_status" gorm:"column:node_status;comment:节点状态 0:未使用 1:使用中 2:异常"`
	// 节点模式
	NodeModel string `json:"nodeModel" form:"node_model" gorm:"column:node_model;comment:节点模式"`
	// GPU类型
	GpuModel string `json:"gpuModel" form:"gpu_model" gorm:"column:gpu_model;comment:GPU类型"`
}

// TableName 节点表名
func (ShareNode) TableName() string {
	return "share_node"
}

// NodeUserRelation 节点用户关联表
type NodeUserRelation struct {
	gorm.Model
	// 用户ID
	UserId string `json:"userId" form:"user_id" gorm:"user_id;comment:用户ID"`
	// 节点PortainerID
	PortainerId string `json:"portainerId" form:"portainer_id" gorm:"portainer_id;comment:PortainerID"`
}

// TableName 节点用户关联表名
func (NodeUserRelation) TableName() string {
	return "node_user_relation"
}

// GpuConRelation gpu容器关联表
type GpuConRelation struct {
	gorm.Model
	// GpuID
	GpuId string `json:"gpuId" gorm:"gpu_id;comment:显卡ID"`
	// 容器ID
	ConId string `json:"conId" gorm:"con_id;comment:容器ID"`
	// 用户ID
	UserId string `json:"userId" gorm:"user_id;comment:用户ID"`
	// 用户名称
	Username string `json:"username" gorm:"username;comment:用户名称"`
}

// TableName gpu容器关联表名
func (GpuConRelation) TableName() string {
	return "gpu_con_relation"
}

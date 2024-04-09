package models

import "gorm.io/gorm"

// GPU监控结构体
type GpuMonitor struct {
	gorm.Model
	// GPU的UUID
	Uuid string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:GPU的UUID;"`
	// GPU卡槽号
	Size string `json:"size" form:"size" gorm:"column:size;comment:GPU卡槽号;"`
	// 所属节点PortainerID
	NodeId string `json:"nodeId" form:"node_id" gorm:"column:node_id;comment:所属节点PortainerID"`
	// 所属节点IP地址
	NodeIpAddress string `json:"nodeIpAddress" form:"node_ip_address" gorm:"column:node_ip_address;comment:所属节点IP地址"`
	// 使用率
	UsageRate string `json:"usageRate" form:"usage_rate" gorm:"column:usage_rate;comment:使用率"`
	// GPU型号
	Type string `json:"type" form:"type" gorm:"column:type;comment:GPU型号"`
	// 温度
	Temperature string `json:"temperature" form:"temperature" gorm:"column:temperature;comment:温度"`
	// 显存
	VMemory string `json:"vMemory" form:"v_memory" gorm:"column:v_memory;comment:总显存"`
	// 使用显存
	VMemoryUsage string `json:"vMemoryUsage" form:"v_memory_usage" gorm:"column:v_memory_usage;comment:使用显存"`
	// 最大功率
	Power string `json:"power" form:"power" gorm:"column:power;comment:最大功率"`
	// 当前功率
	PowerUsage string `json:"powerUsage" form:"power_usage" gorm:"column:power_usage;comment:当前功率"`
	// 功率使用率
	PowerUsageRate string `json:"powerUsageRate" form:"power_usage_rate" gorm:"column:power_usage_rate;comment:功率使用率"`
	// 是否使用
	Used bool `json:"used" form:"used" gorm:"default:false;column:used;comment:是否使用"`
	// 用户ID
	UserId string `json:"userId" form:"user_id" gorm:"column:user_id;comment:用户UUID"`
	// 用户名称
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户名称"`
}

// GPU监控表名
func (GpuMonitor) TableName() string {
	return "gpu_monitor"
}

// ShareGpuMonitor GPU监控结构体
type ShareGpuMonitor struct {
	gorm.Model
	// GPU的UUID
	Uuid string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:GPU的UUID;"`
	// GPU卡槽号
	Size string `json:"size" form:"size" gorm:"column:size;comment:GPU卡槽号;"`
	// 所属节点PortainerID
	NodeId string `json:"nodeId" form:"node_id" gorm:"column:node_id;comment:所属节点PortainerID"`
	// 所属节点IP地址
	NodeIpAddress string `json:"nodeIpAddress" form:"node_ip_address" gorm:"column:node_ip_address;comment:所属节点IP地址"`
	// 使用率
	UsageRate string `json:"usageRate" form:"usage_rate" gorm:"column:usage_rate;comment:使用率"`
	// GPU型号
	Type string `json:"type" form:"type" gorm:"column:type;comment:GPU型号"`
	// 温度
	Temperature string `json:"temperature" form:"temperature" gorm:"column:temperature;comment:温度"`
	// 显存
	VMemory string `json:"vMemory" form:"v_memory" gorm:"column:v_memory;comment:总显存"`
	// 使用显存
	VMemoryUsage string `json:"vMemoryUsage" form:"v_memory_usage" gorm:"column:v_memory_usage;comment:使用显存"`
	// 最大功率
	Power string `json:"power" form:"power" gorm:"column:power;comment:最大功率"`
	// 当前功率
	PowerUsage string `json:"powerUsage" form:"power_usage" gorm:"column:power_usage;comment:当前功率"`
	// 功率使用率
	PowerUsageRate string `json:"powerUsageRate" form:"power_usage_rate" gorm:"column:power_usage_rate;comment:功率使用率"`
	// 是否使用
	Used bool `json:"used" form:"used" gorm:"default:false;column:used;comment:是否使用"`
	// 容器个数
	ConNum uint32 `json:"conNum" gorm:"column:con_num;comment:容器个数;default:0"`
}

// GPU监控表名
func (ShareGpuMonitor) TableName() string {
	return "share_gpu_monitor"
}

// 节点监控结构体
type NodeMonitor struct {
	gorm.Model
	// 节点PortainerID
	PortainerId string `json:"portainerId" form:"portainer_id" gorm:"portainer_id;comment:PortainerID"`
	// IP地址
	IPAddress string `json:"ipAddress" form:"ip_address" gorm:"column:ip_address;comment:IP地址"`
	// GPU类型
	GpuType string `json:"gpuType" form:"gpu_type" gorm:"column:gpu_type;comment:GPU类型"`
	// GPU总量
	GpuQuantity string `json:"gpuQuantity" form:"gpu_quantity" gorm:"gpu_quantity;comment:GPU总数"`
	// GPU使用数量
	GpuUseQuantity string `json:"gpuUseQuantity" form:"gpu_use_quantity" gorm:"gpu_use_quantity;comment:GPU使用数量"`
	// 容器个数
	ContainerQuantity string `json:"containerQuantity" form:"container_quantity" gorm:"column:container_quantity;comment:容器个数"`
	// CPU
	CPUMonitor
	// Memory
	MemoryMonitor
	// Disk
	DiskMonitor
}

// 节点监控表名
func (NodeMonitor) TableName() string {
	return "node_monitor"
}

// 共享节点监控结构体
type ShareNodeMonitor struct {
	gorm.Model
	// 节点PortainerID
	PortainerId string `json:"portainerId" form:"portainer_id" gorm:"portainer_id;comment:PortainerID"`
	// IP地址
	IPAddress string `json:"ipAddress" form:"ip_address" gorm:"column:ip_address;comment:IP地址"`
	// GPU类型
	GpuType string `json:"gpuType" form:"gpu_type" gorm:"column:gpu_type;comment:GPU类型"`
	// GPU总量
	GpuQuantity string `json:"gpuQuantity" form:"gpu_quantity" gorm:"gpu_quantity;comment:GPU总数"`
	// GPU使用数量
	GpuUseQuantity string `json:"gpuUseQuantity" form:"gpu_use_quantity" gorm:"gpu_use_quantity;comment:GPU使用数量"`
	// 容器个数
	ContainerQuantity string `json:"containerQuantity" form:"container_quantity" gorm:"column:container_quantity;comment:容器个数"`
	// CPU
	CPUMonitor
	// Memory
	MemoryMonitor
	// Disk
	DiskMonitor
}

// 节点监控表名
func (ShareNodeMonitor) TableName() string {
	return "share_node_monitor"
}

// CPU结构体
type CPUMonitor struct {
	// CPU型号
	CpuType string `json:"cpuType" form:"cpu_type" gorm:"column:cpu_type;comment:CPU型号"`
	// CPU使用率
	CpuUsage string `json:"cpuUsage" form:"cpu_usage" gorm:"column:cpu_usage;comment:CPU使用率"`
	// CPU总数
	CpuQuantity string `json:"cpuQuantity" form:"cpu_quantity" gorm:"cpu_quantity;comment:CPU总数"`
	// CPU使用数量
	CpuUseQuantity string `json:"cpuUseQuantity" form:"cpu_use_quantity" gorm:"cpu_use_quantity;comment:CPU使用数量"`
}

// Memory结构体
type MemoryMonitor struct {
	// 内存
	Memory string `json:"memory" form:"memory" gorm:"column:memory;comment:内存"`
	// 使用内存
	MemoryUsage string `json:"memoryUsage" form:"memory_usage" gorm:"column:memory_usage;comment:使用内存"`
	// 内存使用率
	MemoryUsageRate string `json:"memoryUsageRate" form:"memory_usage_rate" gorm:"column:memory_usage_rate;comment:内存使用率"`
}

// Disk结构体
type DiskMonitor struct {
	// 硬盘大小
	Disk string `json:"disk" form:"disk" gorm:"column:disk;comment:硬盘大小"`
	// 硬盘使用量
	DiskUsage string `json:"diskUsage" form:"disk_usage" gorm:"column:disk_usage;comment:硬盘使用量"`
	// 硬盘使用率
	DiskUsageRate string `json:"diskUsageRate" form:"disk_usage_rate" gorm:"disk_usage_rate;comment:硬盘使用率"`
}

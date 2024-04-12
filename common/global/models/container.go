package models

import (
	"time"

	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	ContainerId    string `json:"containerId" form:"containerId" gorm:"column:container_id;comment:容器ID;"`
	ContaineName   string `json:"containeName" form:"containeName" gorm:"column:containe_name;comment:容器名称;"`
	PortainerName  string `json:"portainerName" form:"portainerName" gorm:"column:portainer_name;comment:portainer容器名称;"`
	Image          string `json:"image" form:"image" gorm:"column:image;comment:容器镜像;"`
	NodeId         string `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:GPU物理机_PortainerID;"`
	Status         string `json:"status" form:"status" gorm:"column:status;comment:容器状态;"`
	NickName       string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:创建人名称;"`
	UserUuid       string `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:用户UUID;"`
	NodeIp         string `json:"node_ip" form:"node_ip" gorm:"column:node_ip;comment:节点ip;"`
	PublicIp       string `json:"public_ip" form:"public_ip" gorm:"column:public_ip;comment:容器的公网ip;"`
	Password       string `json:"password" form:"password" gorm:"column:password;comment:容器密码;"`
	Gpus           string `json:"gpus" form:"gpus" gorm:"type:text;column:gpus;comment:GPU占用数;"`
	DataVolumeName string `json:"dataVolumeName,optional" form:"dataVolumeName" gorm:"column:data_volume_name;comment:数据卷名称;"`
	OfficialImage  bool   `json:"officialImage,optional,default=true" form:"officialImage" gorm:"-"`
	SharedStorage  bool   `json:"sharedStorage,optional,default=true" form:"sharedStorage" gorm:"column:shared_storage;comment:是否有共享存储;"`
	SubscribeType  int    `json:"subscribeType,optional,default=1" form:"subscribeType" gorm:"column:subscribe_type;comment:订阅类型(1、按时计费2、包时段制);"`
	Expired        bool   `json:"expired,optional" form:"expired" gorm:"column:expired;comment:是否过期;"`
	CheckGpu       bool   `json:"checkGpu,optional" form:"checkGpu" gorm:"column:check_gpu;comment:包时段制续费后,是否检测gpu;"`
	PortMap        string `json:"portMap,optional" form:"portMap" gorm:"column:port_map;comment:端口绑定;"`
	Option         string `json:"option,optional" form:"option" gorm:"column:option;comment:容器配置;"`
	ConvertStatus  int    `json:"convertStatus,optional" form:"convertStatus" gorm:"column:convert_status;comment:转换状态;"`
}

type ContainerSearch struct {
	StartCreatedAt string   `json:"startCreatedAt,optional" form:"startCreatedAt,optional"`
	EndCreatedAt   string   `json:"endCreatedAt,optional" form:"endCreatedAt,optional"`
	Pageinfo       PageInfo `json:"pageinfo,optional"`
}

type GetContainerListResp struct {
	ContainerList []Container `json:"containerList"`
	Total         int64       `json:"total"`
	Page          int         `json:"page"`
	PageSize      int         `json:"pageSize"`
}

type RenameReq struct {
	EndpointId  int32  `json:"endpointId"`
	ContainerId string `json:"containerId"`
	Name        string `json:"name"`
}

type InspContainerReq struct {
	Node        string `json:"node"`
	ContainerId string `json:"containerId"`
}

type CreateExecReq struct {
	EndpointId  int32  `json:"endpointId"`
	ContainerId string `json:"containerId"`
	Cmd         string `json:"cmd"`
}

// 创建exec返回的信息
type CreateExecRsp struct {
	Token  string `json:"token"`
	ExecId string `json:"execId"`
}

type CreateExecBody struct {
	ID           string   `json:"id"`
	AttachStdin  bool     `json:"AttachStdin"`
	AttachStdout bool     `json:"AttachStdout"`
	AttachStderr bool     `json:"AttachStderr"`
	Tty          bool     `json:"Tty"`
	Cmd          []string `json:"Cmd"`
}

type RestartPolicyReq struct {
	Name string `json:"name"`
}

// Name可选值 None / On Failure / Always / Unless Stopped
type UpdateReq struct {
	RestartPolicy RestartPolicyReq `json:"restartPolicy"`
}

type ContainerReq struct {
	EndpointId int32    `json:"endpointId"`
	Ids        []string `json:"ids"`
}

type DeviceRequests struct {
	Driver       string     `json:"driver,optional"`
	Count        int        `json:"count,optional"`
	DeviceIDs    []string   `json:"deviceIDs,optional"`
	Capabilities [][]string `json:"capabilities,optional"`
}

type TarServer struct {
	HostPort string `json:"HostPort,optional"`
}

type ContainerLogReq struct {
	Node      string `json:"node,optional" form:"node,optional"` // 节点名称
	Id        string `json:"id,optional" form:"id,optional"`     // 容器id
	Follow    bool   `json:"follow,default=false,optional"`      // 返回日志后是否继续监听 默认false
	Stdout    bool   `json:"stdout,default=true,optional"`       // 是否返回标准输出 默认true
	Stderr    bool   `json:"stderr,default=true,optional"`       // 是否返回标准错误 默认true
	Since     int    `json:"since,default=0,optional"`           // 返回日志的起始时间 默认0
	Utils     int    `json:"utils,default=0,optional"`           // 返回日志的结束时间 默认0
	Timestamp bool   `json:"timestamp,default=false,optional"`   // 是否显示时间戳 默认false
	Tail      string `json:"tail,default=100,optional"`          // 返回日志的最后几行 从最后一行开始 默认100
}

type ContainerPortReq struct {
	Node int32  `json:"node" form:"node"` // 节点名称
	Id   string `json:"id" form:"id"`     // 容器id
}

type GetContainerPortReq struct {
	Node int32  `json:"node" form:"node"` // 节点名称
	Id   string `json:"id" form:"id"`     // 容器id
	Port int    `json:"port" form:"port"` // 端口类型
}

type SaveImageReq struct {
	Node      int32  `json:"node" form:"node"`           // 节点名称
	Id        string `json:"id" form:"id"`               // 容器id
	ImageName string `json:"imageName" form:"imageName"` // 镜像名称
}

type SingleContainerReq struct {
	Node int32  `json:"node" form:"node"` // 节点名称
	Id   string `json:"id" form:"id"`     // 容器id
}

type CMSContainersReq struct {
	Node           int32      `json:"node" form:"node"` // 节点名称
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	PageInfo
}

type ConvertContainerV21Req struct {
	Id     string `json:"id" form:"id"`         // 容器id
	GpuNum int    `json:"gpuNum" form:"gpuNum"` // gpu数量
	Size   int    `json:"size" form:"size"`     // 数据卷大小
}

// Key 用户SSH公钥表
type Key struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"column:name;comment:公钥名;"`
	Notes    string `json:"notes" form:"notes" gorm:"column:notes;comment:公钥备注;"`
	Key      string `json:"key" form:"key" gorm:"column:key;comment:SSH公钥;type:text;"`
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户名;"`
	UserUuid string `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:用户UUID;"`
}

type CreateKeyReq struct {
	Name  string `json:"name"`
	Notes string `json:"notes"`
	Key   string `json:"key"`
}

type ContainerPortResp struct {
	ContainerId string  `json:"containerId"`
	PublicIp    string  `json:"public_ip"`
	Password    string  `json:"password"`
	Ports       []Ports `json:"ports"`
}

type Ports struct {
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

type ContainerDiskResp struct {
	Type      string `json:"type"`
	Size      string `json:"size"`
	Used      string `json:"used"`
	Available string `json:"available"`
	UsedRate  string `json:"usedRate"`
	Mount     string `json:"mount"`
}

// CMS容器端口信息
type CMSContainerPortResp struct {
	ContainerId string          `json:"container_id"`
	PrivateIp   string          `json:"private_ip"`
	Password    string          `json:"password"`
	Ports       []InternalPorts `json:"ports"`
}

type InternalPorts struct {
	PrivatePort int    `json:"private_port"`
	Type        string `json:"type"`
}

type DeleteContainerReq struct {
	EndpointId int32    `json:"endpointId"`
	Ids        []string `json:"ids"`
}

// 若不设置默认为 containers
func (Container) TableName() string {
	return "container"
}

// 描述容器中需要暴露的端口
type ExposedPorts struct {
	The22TCP   Empty `json:"22/tcp,optional"`
	The5901TCP Empty `json:"5901/tcp,optional"`
	The8888TCP Empty `json:"8888/tcp,optional"`
}

// 容器内部端口与宿主机端口的映射关系
type PortBindings struct {
	The22TCP   []BaseThetype `json:"22/tcp,optional"`
	The5901TCP []BaseThetype `json:"5901/tcp,optional"`
	The8888TCP []BaseThetype `json:"8888/tcp,optional"`
}

type RestartPolicy struct {
	Name string `json:"Name,optional"`
}

// 容器的网络配置
type NetworkingConfig struct {
	EndpointsConfig EndpointsConfig `json:"EndpointsConfig,optional"`
}

type BaseThetype struct {
	HostPort string `json:"HostPort,optional"`
}

// 容器的网络端点配置
type EndpointsConfig struct {
	Bridge Bridge `json:"bridge"`
}

// 指定容器网络的桥接模式
type Bridge struct {
	IPAMConfig IPAMConfig `json:"IPAMConfig,optional"`
}

// 容器的 IP 地址管理配置
type IPAMConfig struct {
	// 表示容器使用IPV4地址
	IPv4Address string `json:"IPv4Address,optional"`
	// 表示容器使用IPV6地址
	IPv6Address string `json:"IPv6Address,optional"`
}

type DeviceRequest struct {
	// 表示设备所需的能力（capabilities）。每个子数组代表一个设备的能力列表，通常由字符串组成。例如，["gpu"] 表示需要 GPU 设备，["net", "inet"] 表示需要网络和互联网访问能力。
	Capabilities [][]string `json:"Capabilities,optional"`
	// 表示需要请求的设备数量
	Count int64 `json:"Count,optional"`
	// 表示请求的设备的唯一标识符列表，每个字符串代表一个设备的标识符，通常是设备名称或ID
	DeviceIDs []string `json:"DeviceIDs,optional"`
	// 表示设备所使用的特定驱动程序的名称，例如 nvidia docker
	Driver string `json:"Driver,optional"`
}

type HostConfig struct {
	// 指定容器是否在停止时删除
	AutoRemove bool `json:"AutoRemove,optional"`
	// 指定容器与宿主机之间的挂载点关系
	Binds []string `json:"Binds,optional"`
	// 指定添加到容器的Linxu内核功能的列表
	CapAdd []string `json:"CapAdd,optional"`
	// 指定从容器中删除的Linux内核功能的列表
	CapDrop []string `json:"CapDrop,optional"`
	// 指定容器对设备的请求
	DeviceRequests []DeviceRequest `json:"DeviceRequests,optional"`
	// 指定容器可以访问的设备列表
	Devices []string `json:"Devices,optional"`
	// 指定容器使用的DNS服务器列表
	DNS []string `json:"Dns,optional"`
	// 指定容器的额外主机名和IP地址
	ExtraHosts []string `json:"ExtraHosts,optional"`
	// 指定是否为容器启用init进程
	Init bool `json:"Init,optional"`
	// 指定容器的内存限制
	Memory int64 `json:"Memory,optional"`
	// 指定容器的内存预留量
	MemoryReservation int64 `json:"MemoryReservation,optional"`
	// 指定容器的CPU资源限制,以纳秒为单位
	NanoCpus int64 `json:"NanoCpus,optional"`
	// 指定容器使用的网络模式
	NetworkMode string `json:"NetworkMode,optional"`
	// 指定容器的端口绑定关系
	PortBindings PortBindings `json:"PortBindings,optional"`
	// 指定是否将容器设置为特权容器
	Privileged bool `json:"Privileged,optional"`
	// 指定是否为容器发布所有端口到主机上
	PublishAllPorts bool `json:"PublishAllPorts,optional"`
	// 指定容器重启策略
	RestartPolicy RestartPolicy `json:"RestartPolicy,optional"`
	// 指定容器的运行时
	Runtime string `json:"Runtime,optional,omitempty"`
	// 指定容器共享内存大小,以字节为单位
	ShmSize int64 `json:"ShmSize,optional"`
	// 指定容器的内核参数
	Sysctls Empty `json:"Sysctls,optional"`
}

type CreateContainerReq struct {
	//镜像 需要在容器中运行的镜像名称
	Image string `json:"image,optional"`
	//环境变量 一个字符串切片 包含要设置在容器中的环境变量
	Env []string `json:"env,optional"`
	//MAC地址 在容器内设置指定的MAC地址
	MACAddress string `json:"macAddress,optional"`
	//暴露端口 指定容器中哪些端口需要暴露以供外部访问
	ExposedPorts ExposedPorts `json:"exposedPorts,optional"`
	//入口点 这是容器启动后要执行的第一个命令
	Entrypoint string `json:"entrypoint,optional,omitempty"`
	//主机配置 包含关于容器如何在主机上运行的配置，如挂载的卷、资源限制等
	HostConfig HostConfig `json:"hostConfig,optional"`
	//网络配置 指定容器的网络配置，如使用哪个网络、IP地址等
	NetworkingConfig NetworkingConfig `json:"networkingConfig,optional"`
	//标签 用于设置容器的标签，可根据标签对容器进行分类和组织
	Labels Empty `json:"labels,optional"`
	//容器名称 用于在Portainer中识别容器
	Name string `json:"name"`
	//打开标准输入 指定是否在容器中打开标准输入
	OpenStdin bool `json:"openStdin,optional"`
	//TTY 指定是否分配终端给容器
	TTY bool `json:"tty,optional"`
	// 指定容器挂载的卷
	Volumes Empty `json:"volumes,optional"`
	// 指定容器连接的类型
	Connection string `json:"connection,optional"`
}

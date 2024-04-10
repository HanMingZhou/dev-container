package container

import "time"

type Config struct {
	Host     string
	Port     int
	Schema   string
	User     string
	Password string
	URL      string
}

type Endpoint struct {
	Id        int32    `json:"Id,omitempty"`
	Name      string   `json:"Name,omitempty"`
	URL       string   `json:"URL,omitempty"`
	PublicURL string   `json:"PublicURL,omitempty"`
	GroupID   int32    `json:"GroupID,omitempty"`
	Tags      []string `json:"Tags"`
}

type Image struct {
	Containers  int         `json:"Containers"`
	Created     int         `json:"Created"`
	ID          string      `json:"Id"`
	Labels      interface{} `json:"Labels"`
	ParentID    string      `json:"ParentId"`
	SharedSize  int         `json:"SharedSize"`
	Size        int         `json:"Size"`
	VirtualSize int         `json:"VirtualSize"`
}

type Container struct {
	ID         string            `json:"Id"`
	Names      []string          `json:"Names"`
	Image      string            `json:"Image"`
	ImageID    string            `json:"ImageID"`
	Command    string            `json:"Command"`
	Created    int               `json:"Created"`
	State      string            `json:"State"`
	Status     string            `json:"Status"`
	Ports      []Ports           `json:"Ports"`
	Labels     map[string]string `json:"Labels,omitempty"`
	SizeRw     int               `json:"SizeRw"`
	SizeRootFs int               `json:"SizeRootFs"`
	HostConfig struct {
		NetworkMode string `json:"NetworkMode"`
	} `json:"HostConfig"`
	NetworkSettings struct {
		Networks struct {
			Bridge struct {
				IPAMConfig          interface{} `json:"IPAMConfig"`
				Links               interface{} `json:"Links"`
				Aliases             interface{} `json:"Aliases"`
				NetworkID           string      `json:"NetworkID"`
				EndpointID          string      `json:"EndpointID"`
				Gateway             string      `json:"Gateway"`
				IPAddress           string      `json:"IPAddress"`
				IPPrefixLen         int         `json:"IPPrefixLen"`
				IPv6Gateway         string      `json:"IPv6Gateway"`
				GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
				GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
				MacAddress          string      `json:"MacAddress"`
			} `json:"bridge"`
		} `json:"Networks"`
	} `json:"NetworkSettings"`
	Mounts []struct {
		Name        string `json:"Name"`
		Source      string `json:"Source"`
		Destination string `json:"Destination"`
		Driver      string `json:"Driver"`
		Mode        string `json:"Mode"`
		RW          bool   `json:"RW"`
		Propagation string `json:"Propagation"`
	} `json:"Mounts"`
}

type Ports struct {
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

/**
 * @Author Flamingo
 * @Description //容器创建返回信息
 * @Date 2023/1/30 16:39
 **/
type CreateConRsp struct {
	ID        string `json:"Id"`
	Portainer struct {
		ResourceControl struct {
			ID             int           `json:"Id"`
			ResourceID     string        `json:"ResourceId"`
			SubResourceIds []interface{} `json:"SubResourceIds"`
			Type           int           `json:"Type"`
			UserAccesses   []struct {
				UserID      int `json:"UserId"`
				AccessLevel int `json:"AccessLevel"`
			} `json:"UserAccesses"`
			TeamAccesses       []interface{} `json:"TeamAccesses"`
			Public             bool          `json:"Public"`
			AdministratorsOnly bool          `json:"AdministratorsOnly"`
			System             bool          `json:"System"`
		} `json:"ResourceControl"`
	} `json:"Portainer"`
	Warnings []interface{} `json:"Warnings"`
}

/**
 * @Author Flamingo
 * @Description //protianer exec 返回数据
 * @Date 2023/2/8 15:05
 **/
type ExecConRsp struct {
	ID string `json:"Id"`
}

// 创建exec返回的信息
type CreateExecRsp struct {
	Token  string `json:"token"`
	ExecId string `json:"execId"`
}

/**
 * @Author Wuenze
 * @Description 容器 inspect 返回信息
 * @Date 15:52 2023/3/27
 **/
type Inspect struct {
	AppArmorProfile string   `json:"AppArmorProfile"`
	Args            []string `json:"Args"`
	Config          struct {
		AttachStderr bool                   `json:"AttachStderr"`
		AttachStdin  bool                   `json:"AttachStdin"`
		AttachStdout bool                   `json:"AttachStdout"`
		Cmd          []string               `json:"Cmd"`
		Domainname   string                 `json:"Domainname"`
		Entrypoint   []string               `json:"Entrypoint"`
		Env          []string               `json:"Env"`
		ExposedPorts map[string]interface{} `json:"ExposedPorts"`
		Hostname     string                 `json:"Hostname"`
		Image        string                 `json:"Image"`
		Labels       struct {
			ComDockerComposeConfigHash         string `json:"com.docker.compose.config-hash"`
			ComDockerComposeContainerNumber    string `json:"com.docker.compose.container-number"`
			ComDockerComposeOneoff             string `json:"com.docker.compose.oneoff"`
			ComDockerComposeProject            string `json:"com.docker.compose.project"`
			ComDockerComposeProjectConfigFiles string `json:"com.docker.compose.project.config_files"`
			ComDockerComposeProjectWorkingDir  string `json:"com.docker.compose.project.working_dir"`
			ComDockerComposeService            string `json:"com.docker.compose.service"`
			ComDockerComposeVersion            string `json:"com.docker.compose.version"`
			Maintainer                         string `json:"maintainer"`
		} `json:"Labels"`
		OnBuild    interface{} `json:"OnBuild"`
		OpenStdin  bool        `json:"OpenStdin"`
		StdinOnce  bool        `json:"StdinOnce"`
		Tty        bool        `json:"Tty"`
		User       string      `json:"User"`
		Volumes    struct{}    `json:"Volumes"`
		WorkingDir string      `json:"WorkingDir"`
	} `json:"Config"`
	Created     time.Time   `json:"Created"`
	Driver      string      `json:"Driver"`
	ExecIDs     interface{} `json:"ExecIDs"`
	GraphDriver struct {
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string `json:"Name"`
	} `json:"GraphDriver"`
	HostConfig struct {
		AutoRemove           bool        `json:"AutoRemove"`
		Binds                []string    `json:"Binds"`
		BlkioDeviceReadBps   interface{} `json:"BlkioDeviceReadBps"`
		BlkioDeviceReadIOps  interface{} `json:"BlkioDeviceReadIOps"`
		BlkioDeviceWriteBps  interface{} `json:"BlkioDeviceWriteBps"`
		BlkioDeviceWriteIOps interface{} `json:"BlkioDeviceWriteIOps"`
		BlkioWeight          int         `json:"BlkioWeight"`
		BlkioWeightDevice    interface{} `json:"BlkioWeightDevice"`
		CapAdd               interface{} `json:"CapAdd"`
		CapDrop              interface{} `json:"CapDrop"`
		Cgroup               string      `json:"Cgroup"`
		CgroupParent         string      `json:"CgroupParent"`
		CgroupnsMode         string      `json:"CgroupnsMode"`
		ConsoleSize          []int       `json:"ConsoleSize"`
		ContainerIDFile      string      `json:"ContainerIDFile"`
		CpuCount             int         `json:"CpuCount"`
		CpuPercent           int         `json:"CpuPercent"`
		CpuPeriod            int         `json:"CpuPeriod"`
		CpuQuota             int         `json:"CpuQuota"`
		CpuRealtimePeriod    int         `json:"CpuRealtimePeriod"`
		CpuRealtimeRuntime   int         `json:"CpuRealtimeRuntime"`
		CpuShares            int         `json:"CpuShares"`
		CpusetCpus           string      `json:"CpusetCpus"`
		CpusetMems           string      `json:"CpusetMems"`
		DeviceCgroupRules    interface{} `json:"DeviceCgroupRules"`
		DeviceRequests       interface{} `json:"DeviceRequests"`
		Devices              interface{} `json:"Devices"`
		Dns                  interface{} `json:"Dns"`
		DnsOptions           interface{} `json:"DnsOptions"`
		DnsSearch            interface{} `json:"DnsSearch"`
		ExtraHosts           interface{} `json:"ExtraHosts"`
		GroupAdd             interface{} `json:"GroupAdd"`
		IOMaximumBandwidth   int         `json:"IOMaximumBandwidth"`
		IOMaximumIOps        int         `json:"IOMaximumIOps"`
		IpcMode              string      `json:"IpcMode"`
		Isolation            string      `json:"Isolation"`
		KernelMemory         int         `json:"KernelMemory"`
		KernelMemoryTCP      int         `json:"KernelMemoryTCP"`
		Links                interface{} `json:"Links"`
		LogConfig            struct {
			Config struct {
			} `json:"Config"`
			Type string `json:"Type"`
		} `json:"LogConfig"`
		MaskedPaths       []string               `json:"MaskedPaths"`
		Memory            int                    `json:"Memory"`
		MemoryReservation int                    `json:"MemoryReservation"`
		MemorySwap        int                    `json:"MemorySwap"`
		MemorySwappiness  interface{}            `json:"MemorySwappiness"`
		NanoCpus          int                    `json:"NanoCpus"`
		NetworkMode       string                 `json:"NetworkMode"`
		OomKillDisable    bool                   `json:"OomKillDisable"`
		OomScoreAdj       int                    `json:"OomScoreAdj"`
		PidMode           string                 `json:"PidMode"`
		PidsLimit         interface{}            `json:"PidsLimit"`
		PortBindings      map[string]interface{} `json:"PortBindings"`
		Privileged        bool                   `json:"Privileged"`
		PublishAllPorts   bool                   `json:"PublishAllPorts"`
		ReadonlyPaths     []string               `json:"ReadonlyPaths"`
		ReadonlyRootfs    bool                   `json:"ReadonlyRootfs"`
		RestartPolicy     struct {
			MaximumRetryCount int    `json:"MaximumRetryCount"`
			Name              string `json:"Name"`
		} `json:"RestartPolicy"`
		Runtime      string        `json:"Runtime"`
		SecurityOpt  interface{}   `json:"SecurityOpt"`
		ShmSize      int           `json:"ShmSize"`
		UTSMode      string        `json:"UTSMode"`
		Ulimits      interface{}   `json:"Ulimits"`
		UsernsMode   string        `json:"UsernsMode"`
		VolumeDriver string        `json:"VolumeDriver"`
		VolumesFrom  []interface{} `json:"VolumesFrom"`
	} `json:"HostConfig"`
	HostnamePath string `json:"HostnamePath"`
	HostsPath    string `json:"HostsPath"`
	Id           string `json:"Id"`
	Image        string `json:"Image"`
	LogPath      string `json:"LogPath"`
	MountLabel   string `json:"MountLabel"`
	Mounts       []struct {
		Destination string `json:"Destination"`
		Mode        string `json:"Mode"`
		Propagation string `json:"Propagation"`
		RW          bool   `json:"RW"`
		Source      string `json:"Source"`
		Type        string `json:"Type"`
		Driver      string `json:"Driver,omitempty"`
		Name        string `json:"Name,omitempty"`
	} `json:"Mounts"`
	Name            string `json:"Name"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		EndpointID             string `json:"EndpointID"`
		Gateway                string `json:"Gateway"`
		GlobalIPv6Address      string `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int    `json:"GlobalIPv6PrefixLen"`
		HairpinMode            bool   `json:"HairpinMode"`
		IPAddress              string `json:"IPAddress"`
		IPPrefixLen            int    `json:"IPPrefixLen"`
		IPv6Gateway            string `json:"IPv6Gateway"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		MacAddress             string `json:"MacAddress"`
		Networks               struct {
			PrometheusDefault struct {
				Aliases             []string    `json:"Aliases"`
				DriverOpts          interface{} `json:"DriverOpts"`
				EndpointID          string      `json:"EndpointID"`
				Gateway             string      `json:"Gateway"`
				GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
				GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
				IPAMConfig          interface{} `json:"IPAMConfig"`
				IPAddress           string      `json:"IPAddress"`
				IPPrefixLen         int         `json:"IPPrefixLen"`
				IPv6Gateway         string      `json:"IPv6Gateway"`
				Links               []string    `json:"Links"`
				MacAddress          string      `json:"MacAddress"`
				NetworkID           string      `json:"NetworkID"`
			} `json:"prometheus_default"`
		} `json:"Networks"`
		Ports                  map[string]interface{} `json:"Ports"`
		SandboxID              string                 `json:"SandboxID"`
		SandboxKey             string                 `json:"SandboxKey"`
		SecondaryIPAddresses   interface{}            `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses interface{}            `json:"SecondaryIPv6Addresses"`
	} `json:"NetworkSettings"`
	Path           string `json:"Path"`
	Platform       string `json:"Platform"`
	ProcessLabel   string `json:"ProcessLabel"`
	ResolvConfPath string `json:"ResolvConfPath"`
	RestartCount   int    `json:"RestartCount"`
	State          struct {
		Dead       bool      `json:"Dead"`
		Error      string    `json:"Error"`
		ExitCode   int       `json:"ExitCode"`
		FinishedAt time.Time `json:"FinishedAt"`
		OOMKilled  bool      `json:"OOMKilled"`
		Paused     bool      `json:"Paused"`
		Pid        int       `json:"Pid"`
		Restarting bool      `json:"Restarting"`
		Running    bool      `json:"Running"`
		StartedAt  time.Time `json:"StartedAt"`
		Status     string    `json:"Status"`
	} `json:"State"`
}

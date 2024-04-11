package container

type ContainerCon struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	Schema      string `mapstructure:"schema" json:"schema" yaml:"schema"`
	User        string `mapstructure:"user" json:"user" yaml:"user"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	URL         string `mapstructure:"url" json:"url" yaml:"url"`
	Token       string `mapstructure:"token" json:"token" yaml:"token"`
	ServiceName string `mapstructure:"service-name" json:"service-name" yaml:"service-name"`
}

type Portainer struct {
	Config    ContainerCon
	Token     string
	AuthToken string
	ApiURL    string
	Addr      string
}

// DOKCER 容器的实例化方法
func NewContainer() (*Portainer, error) {
	client := NewPortainer(&Portainer{
		Config: ContainerCon{
			//Host:     "192.168.0.78",
			Host:     "192.168.0.53",
			Port:     9000,
			Schema:   "http",
			User:     "admin",
			Password: "cDfQtt2FQgvuHXz",
			URL:      "/api",
			Token:    "ptr_FzuDBJ3zMueL7gRbCCGfY7yaavdNOxEIGUdEzJYKGV0=",
		},
	})
	err := client.Auth()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

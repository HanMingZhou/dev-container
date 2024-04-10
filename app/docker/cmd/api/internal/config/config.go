package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Mysql struct {
		DataSource string
	}

	DockerAccount struct {
		UserPasswd string
		ConPrefix  string
	}

	// harbor配置
	Harbor struct {
		Username string
		Password string
		Host     string
		Official string
	}

	AES struct {
		Key string
	}
}

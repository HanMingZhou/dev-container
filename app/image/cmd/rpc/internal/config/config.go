package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Portainer struct {
		Config struct {
			Host        string
			Port        int
			Schema      string
			User        string
			Password    string
			URL         string
			Token       string
			ServiceName string
		}
		Token     string
		AuthToken string
		ApiURL    string
		Addr      string
	}
}

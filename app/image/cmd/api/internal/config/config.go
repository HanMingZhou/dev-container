package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	// Mysql数据库配置
	Mysql struct {
		DataSource string
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
	//rpc 服务
	ImageRpcConf zrpc.RpcClientConf
}

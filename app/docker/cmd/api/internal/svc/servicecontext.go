package svc

import (
	"go-zero-container/app/docker/cmd/api/internal/config"
	"go-zero-container/common/initDB"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	//Portiner *portiner.Portiner
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     initDB.InitGorm(c.Mysql.DataSource),
		//Portiner: portiner.NewPortiner(zrpc.MustNewClient(c.PortinerRpcConf)),
	}
}

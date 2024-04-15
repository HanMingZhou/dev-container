package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-container/app/image/cmd/api/internal/config"
	"go-zero-container/app/image/cmd/rpc/imageserver"
	"go-zero-container/common/initDB"
	"go-zero-container/common/utils/container"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	Portainer *container.Portainer
	ImageRpc  imageserver.ImageServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB:        initDB.InitGorm(c.Mysql.DataSource),
		Portainer: container.NewContainer(),
		ImageRpc:  imageserver.NewImageServer(zrpc.MustNewClient(c.ImageRpcConf)),
	}
}

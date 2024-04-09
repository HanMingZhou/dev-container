package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-container/app/usercerter/cmd/api/internal/config"
	"go-zero-container/app/usercerter/cmd/rpc/usercenter"
	"go-zero-container/common/initDB"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	DB            *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDB.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		DB:            db,
	}
}
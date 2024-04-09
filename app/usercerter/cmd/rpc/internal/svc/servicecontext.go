package svc

import (
	"go-zero-container/app/usercerter/cmd/rpc/internal/config"
	"go-zero-container/common/initDB"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := initDB.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

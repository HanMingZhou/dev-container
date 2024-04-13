package svc

import (
	"go-zero-container/app/docker/cmd/api/internal/config"
	"go-zero-container/common/initDB"
	"go-zero-container/common/utils/container"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Portiner *container.Portainer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		DB:       initDB.InitGorm(c.Mysql.DataSource),
		Portiner: container.NewContainer(),
	}
}

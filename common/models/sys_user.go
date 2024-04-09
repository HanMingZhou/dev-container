package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUser struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"uuid"`
	ID          int64     `gorm:"id"`
	Username    string    `gorm:"username"`
	Password    string    `gorm:"password"`
	NickName    string    `gorm:"nickname"`
	HeaderImg   string    `gorm:"headerimg,optional"`
	AuthorityId uint      `gorm:"authorityid"`
	Enable      int       `gorm:"enable"`
	Phone       string    `gorm:"phone"`
	Email       string    `gorm:"email"`
}

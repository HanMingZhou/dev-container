package models

import "gorm.io/gorm"

type SysAuthority struct {
	gorm.Model
	AuthorityId   uint   `gorm:"authorityid"`
	AuthorityName string `gorm:"authorityname"`
	ParentId      int    `gorm:"parentid"`
	DefaultRouter string `gorm:"defaultrouter,default:dashboard"`
}

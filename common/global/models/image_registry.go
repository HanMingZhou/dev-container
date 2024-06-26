package models

// ImageRegistry 镜像仓库结构体
type ImageRegistry struct {
	//gorm.Model
	Rid            int64  `json:"rid" gorm:"column:rid;comment:portainer仓库id"`
	Kind           int32  `json:"kind" gorm:"column:kind;comment:类型(1公有 2私有);"`
	UserId         uint   `json:"user_id" gorm:"column:user_id;comment:用户id;"`
	Name           string `json:"name" gorm:"column:name;comment:仓库名称;"`
	Url            string `json:"url" gorm:"column:url;comment:仓库地址;"`
	Authentication int32  `json:"authentication" gorm:"column:authentication;comment:是否开启认证(1开启 2不开启);"`
	Username       string `json:"username" gorm:"column:username;comment:用户名;"`
	Password       string `json:"password" gorm:"column:password;comment:密码;"`
}

// 设置默认表为 image_registry
func (ImageRegistry) TableName() string {
	return "image_registry"
}

package initDB

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化gorm连接数据库
func InitGorm(MysqlDataSource string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDataSource), &gorm.Config{})
	if err != nil {
		panic("系统初始化连接数据库失败，error：" + err.Error())
	} else {
		fmt.Println("连接数据库成功")
	}
	return db
}

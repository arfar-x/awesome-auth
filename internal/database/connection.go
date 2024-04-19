package database

import (
	"fmt"

	"awesome-auth/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDbClient(config *configs.AppConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port,
		config.DB.Name, config.DB.Charset,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

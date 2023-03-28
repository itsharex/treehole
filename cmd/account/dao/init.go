package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitDB() {
	dsn := viper.GetString("db.mysql.dsn")
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.Default(), logger.Config{
			IgnoreRecordNotFoundError: true,
		}),
	})
	if err != nil {
		panic(err)
	}
	SetDefault(open)
}

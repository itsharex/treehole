package main

import (
	"github.com/Jazee6/treehole/cmd/account/dao"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := viper.GetString("mysql.dsn.account")
	open, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}
	dao.SetDefault(open)
}

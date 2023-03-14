package main

import (
	_ "github.com/Jazee6/treehole/account/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := viper.GetString("mysql_dsn")

	open, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./pkg/dao",
		ModelPkgPath: "./model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(open)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

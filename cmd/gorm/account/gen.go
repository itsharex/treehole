package main

import (
	"fmt"
	_ "github.com/Jazee6/treehole/pkg/configs"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := viper.GetString("db.mysql.account.dsn")
	fmt.Println(dsn)

	open, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./cmd/account/dao",
		ModelPkgPath:      "./model",
		WithUnitTest:      false,
		FieldNullable:     false,
		FieldCoverable:    false,
		FieldSignable:     true,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(open)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

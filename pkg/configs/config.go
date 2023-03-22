package configs

import "github.com/spf13/viper"

func init() {
	viper.SetConfigFile("config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

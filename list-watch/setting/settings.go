package setting

import (
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./config/")
	config.SetConfigName("config.yaml")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("找不到配置文件..")
		} else {
			panic("配置文件出错..")
		}
	}
	return config
}

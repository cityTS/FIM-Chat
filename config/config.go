package config

import (
	"github.com/FengZhg/go_tools/logger"
	"github.com/spf13/viper"
)

func init() {
	config := viper.New()
	// 初始化
	config.SetDefault("DB.dsn", "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	config.AddConfigPath("./config")
	config.SetConfigType("yml")
	config.SetConfigName("config")

	if err := config.ReadInConfig(); err != nil {
		logger.Log.Panic("read config err: ", err)
	}
	Dsn = config.Get("DB.dsn").(string)

	logger.Log.Info("读取配置文件完成")
}

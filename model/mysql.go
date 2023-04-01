package model

import (
	"FIM-Chat/config"
	"github.com/FengZhg/go_tools/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Panic("false to connect database:", err)
	}
	DB = db
}

// Persistence 消息持久化
func (m *Message) Persistence() bool {
	result := DB.Create(&m)
	if result.Error != nil {
		logger.Log.Warn(result.Error)
		return false
	}
	if result.RowsAffected <= 0 {
		return false
	}
	return true
}

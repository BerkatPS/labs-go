package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig() *Config {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root@tcp(localhost:3306)/chat?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Can't Connect To Database")
	}

	return &Config{DB: db}

}

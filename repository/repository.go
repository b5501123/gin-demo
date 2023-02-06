package repository

import (
	"fmt"
	"gin-demo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func InitDB() {
	dbSetting := &config.Setting.DBConfig
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", dbSetting.User, dbSetting.Password, dbSetting.Host, dbSetting.Port, dbSetting.Database)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

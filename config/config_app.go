package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type AppConfig struct {
	DBConfig  DBConfig
	WebConfig WebConfig
}

type DBConfig struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type WebConfig struct {
	Environment  string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JwtSecret    string
}

var cfg *ini.File

var Setting *AppConfig = &AppConfig{
	DBConfig{}, WebConfig{},
}

func Load() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("web", &Setting.WebConfig)
	mapTo("database", &Setting.DBConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

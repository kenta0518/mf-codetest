package infra

import (
	"github.com/kenta0518/mf-codetest/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlConnector struct {
	DB *gorm.DB
}

func NewMySqlConnector(cfg *config.Config) *MySqlConnector {
	conn := cfg.MySQL.DBConn

	log := logger.Default
	if cfg.IsDevelopment() {
		log = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: log,
	})

	if err != nil {
		panic(err)
	}

	return &MySqlConnector{
		DB: db,
	}
}

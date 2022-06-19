package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Dsn string = "user:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func GenerateDB(dsn string) error {
	var err error

	DB, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(fmt.Errorf("connect DB failed: %w", err))
	}

	return err
}

package config

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/cms?charset=utf8mb4&parseTime=True&loc=Local", "root", "Qwerty@123")
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil

}

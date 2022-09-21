package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	user := GetEnv("MYSQL_USER")
	password := GetEnv("MYSQL_PASSWORD")
	database := GetEnv("MYSQL_DATABASE")
	host := GetEnv("MYSQL_HOST")
	port := GetEnv("MYSQL_PORT")
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	DB = db
	return err
}

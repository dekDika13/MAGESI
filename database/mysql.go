package database

import (
	"fmt"
	"magesi/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Name     string
	DB_Address  string
}

func InitDB() {
	config := Config{
		DB_Username: os.Getenv("DB_USERNAME"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Address:  os.Getenv("DB_ADDRESS"),
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DB_Username, config.DB_Password, config.DB_Address, config.DB_Name)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString))
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	_ = DB.AutoMigrate(&model.Roles{}, &model.Users{})
}

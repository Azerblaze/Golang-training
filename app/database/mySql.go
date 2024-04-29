package database

import (
	"fmt"
	"training/app/configs"
	"training/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	DBUser := configs.Cfg.DBUser
	DBPassword := configs.Cfg.DBPassword
	DBHost := configs.Cfg.DBHost
	DBPort := configs.Cfg.DBPort
	DBName := configs.Cfg.DBName

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Print("halo")
		panic(err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		fmt.Print("hai")
		panic(err)
	}
}

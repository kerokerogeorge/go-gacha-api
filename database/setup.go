package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	godotenv "github.com/joho/godotenv"
)

var DB *gorm.DB

func DbConnect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to read file: ", err)
	}
	MYSQL := "mysql"
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := "tcp(db:3306)"
	DBNAME := os.Getenv("DB_DATABASE")
	CONNECT := USER + ":" + PASS + "@" + HOST + "/" + DBNAME + "?parseTime=true"

	database, err := gorm.Open(MYSQL, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	// database.AutoMigrate(&model.User{})
	// database.DropTable(&model.User{})
	// database.DropTable(&model.Character{})
	// database.AutoMigrate(&model.Result{})
	// database.AutoMigrate(&model.Character{})
	// database.AutoMigrate(&model.Result{})
	// database.AutoMigrate(&model.Gacha{})
	// database.AutoMigrate(&model.CharacterEmmitionRate{})
	// database.DropTable(&model.CharacterEmmitionRate{})
	// database.DropTable(&model.Gacha{})
	// database.DropTable(&model.Gacha{})

	DB = database

	return DB
}

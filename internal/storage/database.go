package storage

import (
	"github.com/blenz/user-service/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

func InitDb() *gorm.DB {

	db, err := gorm.Open("postgres", "dbname=users sslmode=disable")
	if err != nil {
		log.Fatalln(err.Error())
	}

	initTables(db)

	return db
}

func initTables(db *gorm.DB) {
	
	if db.HasTable(&models.User{}) {
		db.DropTableIfExists(&models.User{})
	}
	db.CreateTable(&models.User{})

	// TODO: clean mock data
	db.Create(&models.User{Username: "User1"})
	db.Create(&models.User{Username: "User2"})
	db.Create(&models.User{Username: "User3"})
	db.Create(&models.User{Username: "User4"})
}

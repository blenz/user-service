package storage

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/blenz/user-service/internal/models"
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
}
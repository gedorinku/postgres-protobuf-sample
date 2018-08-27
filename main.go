package main

import (
	"github.com/gedorinku/postgres-protobuf-sample/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	c, err := loadConfig()
	if err != nil {
		panic(err)
	}

	db := connectDB(c.DataBaseURL)
	err = initDB(db)
	if err != nil {
		panic(err)
	}

	if err := sample(db); err != nil {
		panic(err)
	}
}

func connectDB(url string) *gorm.DB {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return db
}

func initDB(db *gorm.DB) error {
	return db.AutoMigrate(models.Book{}).Error
}

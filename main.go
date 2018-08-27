package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	c, err := loadConfig()
	if err != nil {
		panic(err)
	}

	connectDB(c.DataBaseURL)
}

func connectDB(url string) *gorm.DB {
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	return db
}

func initDB(db *gorm.DB) {

}

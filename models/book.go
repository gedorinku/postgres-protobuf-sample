package main

type Book struct {
	ID      uint `gorm:"primary_key"`
	Details []byte
}

package main

import (
	"fmt"

	"github.com/gedorinku/postgres-protobuf-sample/models"
	"github.com/gogo/protobuf/proto"
	"github.com/jinzhu/gorm"
)

var sampleBooks = []*Book{
	{
		Title: "hogeほげ",
		Authors: []*Author{
			{
				Name: "gedorinku",
			},
			{
				Name: "ぴよぴよ",
			},
		},
		PageCount: 100,
	},
	{
		Title:     "ぽよ",
		Authors:   []*Author{},
		PageCount: 300,
	},
}

func sample(db *gorm.DB) error {
	books, err := write(db)
	if err != nil {
		return err
	}

	return read(db, books)
}

func write(db *gorm.DB) ([]*models.Book, error) {
	books := []*models.Book{}
	for _, b := range sampleBooks {
		out, err := proto.Marshal(b)
		if err != nil {
			return nil, err
		}

		book := &models.Book{
			Details: out,
		}

		if err := db.Create(book).Error; err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func read(db *gorm.DB, books []*models.Book) error {
	for _, b := range books {
		if err := db.Model(models.Book{ID: b.ID}).First(b).Error; err != nil {
			return err
		}

		book := &Book{}
		if err := proto.Unmarshal(b.Details, book); err != nil {
			return err
		}

		fmt.Printf("%+v\n", book)
		fmt.Printf("%v", book.Title)
	}

	return nil
}

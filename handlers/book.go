package handlers

import (
	"fmt"

	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/database"
	"github.com/cumhurmesuttokalak/go-bitirmeprojesi/models"
)

type Book struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	NumberOfPages int     `json:"numberOfPages"`
	Publisher     string  `json:"publisher"`
	Price         float64 `json:"price"`
}

func CreateResponseBook(modelBook *models.Book) Book {
	return Book{ID: modelBook.ID, Name: modelBook.Name, Author: modelBook.Author, NumberOfPages: modelBook.NumberOfPages, Publisher: modelBook.Publisher, Price: modelBook.Price}
}
func FindBook(id int, book *models.Book) error {
	database.Database.Db.Find(&book, "id", id)
	if book.ID == 0 {
		fmt.Println("book does not exist")
	}
	return nil
}

//TODO: book oparations

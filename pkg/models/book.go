package models

import (
	"errors"

	"github.com/dp487/scaling-waddle/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

// Book represents a book entity in the database
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init initializes the database connection and auto-migrates the schema
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook creates a new book record in the database
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookById retrieves a book by its ID
func GetBookById(Id int64) (*Book, error) {
	var getBook Book
	result := db.Where("ID = ?", Id).First(&getBook)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}
	return &getBook, nil
}

// UpdateBook updates an existing book record in the database
func UpdateBook(b *Book) error {
	result := db.Save(b)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteBook deletes a book record by its ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}

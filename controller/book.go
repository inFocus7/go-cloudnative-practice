package controller

import (
	"fabiangonz98/go-cloudnative-practice/model"
	"github.com/jinzhu/gorm"
)

func ListBooks(db *gorm.DB) (model.Books, error) {
	books := make([]*model.Book, 0)
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func ReadBook(db *gorm.DB, id uint) (*model.Book, error) {
	book := &model.Book{}
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(db *gorm.DB, id uint) error {
	book := &model.Book{}
	if err := db.Delete(&book, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBook(db *gorm.DB, book *model.Book) error {
	if err := db.First(&model.Book{}, book.ID).Update(book).Error; err != nil {
		return err
	}

	return nil
}

func CreateBook(db *gorm.DB, book *model.Book) (*model.Book, error) {
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

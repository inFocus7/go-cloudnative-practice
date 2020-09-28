package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Books []*Book

// Maps DB records.
type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
}

type BooksView []*BookView

// The shown inner program view/version of a book.
type BookView struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

func (book Book) ToView() *BookView {
	return &BookView{
		book.ID,
		book.Title,
		book.Author,
		book.PublishedDate.Format("2006-01-02"),
		book.ImageURL,
		book.Description,
	}
}

func (books Books) ToView() BooksView {
	res := make([]*BookView, len(books))
	for i, b := range books {
		res[i] = b.ToView()
	}

	return res
}

type BookForm struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

func (form *BookForm) ToModel() (*Book, error) {
	pubDate, err := time.Parse("2006-01-02", form.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &Book{
		Title:         form.Title,
		Author:        form.Author,
		PublishedDate: pubDate,
		ImageURL:      form.ImageURL,
		Description:   form.Description,
	}, nil
}

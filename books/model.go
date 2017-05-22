package books

import (
	"errors"
	"fmt"

	"github.com/didiyudha/go-mongo/config"
	"gopkg.in/mgo.v2/bson"
)

// Book Model
type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float64
}

// AllBooks return all books
func AllBooks() ([]Book, error) {
	books := []Book{}
	err := config.Books.Find(bson.M{}).All(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// OneBook get book by isbn
func OneBook(isbn string) (Book, error) {
	book := Book{}
	if isbn == "" {
		return book, errors.New("ISBN can not be empty or nil")
	}
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&book)
	if err != nil {
		return book, err
	}
	return book, nil
}

// SaveBook save a book
func SaveBook(book Book) (Book, error) {
	err := check(book)
	if err != nil {
		return book, err
	}
	err = config.Books.Insert(book)
	if err != nil {
		return book, err
	}
	return book, nil
}

// UpdateBook updating book record
func UpdateBook(isbn string, book Book) (Book, error) {
	if isbn == "" {
		return book, errors.New("Please send isbn parameters")
	}
	fmt.Println("Pass isbn checking, isbn: ", isbn)
	err := check(book)
	if err != nil {
		return book, err
	}
	fmt.Println("Pass all constraints checking")
	_, err = OneBook(isbn)
	if err != nil {
		fmt.Println("Isbn not found")
		return book, err
	}
	fmt.Println("Isbn found")
	err = config.Books.Update(bson.M{"isbn": isbn}, &book)
	if err != nil {
		return book, err
	}
	fmt.Println("Update success, updated book: ", book)
	return book, nil
}

// DeleteBook delete book by isbn
func DeleteBook(isbn string) error {
	if isbn == "" {
		return errors.New("Please send isbn paramter")
	}
	_, err := OneBook(isbn)
	if err != nil {
		return errors.New("Book with " + isbn + " can not be found")
	}
	err = config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return err
	}
	return nil
}

func check(book Book) error {
	if book.Author == "" || book.Isbn == "" || book.Title == "" {
		return errors.New("Please Complete Author, Isbn, Title or Price field")
	}
	return nil
}

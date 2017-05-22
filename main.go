package main

import (
	"net/http"

	"github.com/didiyudha/go-mongo/books"
)

func main() {
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/show", books.Show)
	http.HandleFunc("/books/new", books.NewBook)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/edit", books.Edit)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/delete", books.Delete)
}

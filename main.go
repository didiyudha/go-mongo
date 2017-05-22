package main

import (
	"fmt"
	"net/http"

	"github.com/didiyudha/go-mongo/books"
)

func main() {
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/books/show", books.Show)
	http.HandleFunc("/books/new", books.NewBook)
	http.HandleFunc("/books/create", books.Create)
	http.HandleFunc("/books/edit", books.Edit)
	http.HandleFunc("/books/update", books.Update)
	http.HandleFunc("/books/delete", books.Delete)
	fmt.Println("Server is run on port 8080")
	http.ListenAndServe(":8080", nil)
}

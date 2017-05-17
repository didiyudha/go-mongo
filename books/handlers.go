package books

import (
	"net/http"
	"strconv"

	"github.com/didiyudha/go-mongo/config"
)

// Index Page
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	books, err := AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	config.TPL.ExecuteTemplate(w, "books.gohtml", books)
}

// Show a single book by isbn
func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")
	book, err := OneBook(isbn)
	if err != nil {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	config.TPL.ExecuteTemplate(w, "show.gohtml", book)
}

// NewBook return a form to create a book
func NewBook(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "create.gohtml", nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Price must be a number", http.StatusBadRequest)
		return
	}
	book := Book{
		Isbn:   r.FormValue("isbn"),
		Author: r.FormValue("author"),
		Title:  r.FormValue("title"),
		Price:  price,
	}
	book, err = SaveBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	config.TPL.ExecuteTemplate(w, "created.gohtml", book)
}

// Edit get book by isbn and showing to form
func Edit(w http.ResponseWriter, r *http.Request) {
	isbn := r.FormValue("isbn")
	book, err := OneBook(isbn)
	if err != nil {
		http.Error(w, "ISBN can not be found", http.StatusNoContent)
		return
	}
	config.TPL.ExecuteTemplate(w, "edit.gohtml", book)
}

// Update data book
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Price must be a number", http.StatusBadRequest)
		return
	}
	isbn := r.FormValue("isbn")
	book := Book{
		Isbn:   isbn,
		Author: r.FormValue("author"),
		Title:  r.FormValue("title"),
		Price:  price,
	}
	book, err = UpdateBook(isbn, book)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	config.TPL.ExecuteTemplate(w, "updated.gohtml", book)
}

// Delete a book
func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, "Please send isbn", http.StatusBadRequest)
		return
	}
	err := DeleteBook(isbn)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/books", http.StatusSeeOther)
}

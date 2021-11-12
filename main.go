package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/muhammad-rz/bookstore/models"
	"log"
	"net/http"
)

type App struct {
	books models.BookModel
}

func main() {
	var err error

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	app := &App{
		books: models.BookModel{DB: db},
	}

	http.HandleFunc("/books", app.booksIndex)
	http.ListenAndServe(":3000", nil)

}


func(app *App) booksIndex(w http.ResponseWriter, r *http.Request) {
	bks, err := app.books.FindAll()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}

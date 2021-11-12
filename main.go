package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/muhammad-rz/bookstore/models"
	"log"
	"net/http"
)

func injectDB(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "db", db)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func main(){
	var err error

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/books", injectDB(db, booksIndex))
	http.ListenAndServe(":3000", nil)

}

func booksIndex(w http.ResponseWriter, r *http.Request){
	bks, err := models.AllBooks(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}


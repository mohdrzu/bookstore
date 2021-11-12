package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/muhammad-rz/bookstore/models"
	"log"
	"net/http"
)

func main(){
	err := models.InitDB("postgres://postgres:postgres@localhost/bookstore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/books", booksIndex)
	http.ListenAndServe(":3000", nil)

}

func booksIndex(w http.ResponseWriter, r *http.Request){
	bks, err := models.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, £%.2f\n", bk.Isbn, bk.Title, bk.Author, bk.Price)
	}
}


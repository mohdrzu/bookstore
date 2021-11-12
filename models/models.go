package models

import "database/sql"

type Book struct {
	Isbn string
	Title string
	Author string
	Price float32
}

type BookModel struct {
	DB *sql.DB
}

func(m BookModel) FindAll() ([]Book, error){
	rows, err := m.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []Book
	for rows.Next() {
		var bk Book
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bks, nil
}
package repository

import (
	"fmt"
	"wc/model"

	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetBooksByAuthor(authorID int) ([]model.Book, error) {
	query := fmt.Sprintf(`SELECT b.id, b.title, b.author_id, b.cost FROM %s b INNER JOIN %s a on a.id = b.author_id
	WHERE a.id = $1`, "bookTable", "authorTable")

	var books []model.Book
	rows, err := r.db.Query(query, authorID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Id, &book.Title, &book.AuthorId, &book.Cost)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookPostgres) GetAuthorBooksCount(authorID int) (int, error) {
	var booksCount int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s b INNER JOIN %s a on a.id = b.author_id
	WHERE a.id=$1`, "bookTable", "authorTable")
	err := r.db.QueryRow(query, authorID).Scan(&booksCount)
	return booksCount, err
}

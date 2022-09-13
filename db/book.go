package db

import (
	"context"
	"database/sql"
	"math/big"
)

const (
	createBookQuery = `INSERT INTO book (
		id,
		bookName,
		description,
		totalCopies,
		currentCopies	
		)
        VALUES(?,?,?,?,?)`

	listBookQuery       = `SELECT * FROM book`
	findBookByIDQuery   = `SELECT * FROM book WHERE id = ?`
	deleteBookByIDQuery = `DELETE FROM book WHERE id = ?`
	updateBookQuery     = `UPDATE book SET bookName=? ,
	description=?, totalCopies=?, currentCopies=? where id = ?`
)

type Book struct {
	ID            string `db:"id"`
	BookName      string `db:"bookName"`
	Description   string `db:"description"`
	TotalCopies   int    `db:"totalCopies"`
	CurrentCopies int    `db:"currentCopies"`
}

type BookItem struct {
	ID          string  `db:"id"`
	Price       int     `db:"price"`
	BookName    string  `db:"bookName"`
	BookId      string  `db:bookId"`
	Borrowed_at big.Int `db:"borrowed_at"`
	Due_date    big.Int `db:"due_date"`
	Return_date big.Int `db:return_date"`
	Lend_by     string  `db:"lend_by"`
}

func (s *store) CreateBook(ctx context.Context, book *Book) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createBookQuery,
			book.ID,
			book.BookName,
			book.Description,
			book.TotalCopies,
			book.CurrentCopies,
		)
		return err
	})
}

func (s *store) ListBooks(ctx context.Context) (books []Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &books, listBookQuery)
	})
	if err == sql.ErrNoRows {
		return books, ErrBookNotExist
	}
	return
}

func (s *store) FindBookByID(ctx context.Context, id string) (book Book, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &book, findBookByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return book, ErrBookNotExist
	}
	return
}

func (s *store) DeleteBookByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteBookByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrBookNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateBook(ctx context.Context, book *Book) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateBookQuery,
			book.BookName,
			book.Description,
			book.TotalCopies,
			book.CurrentCopies,
			book.ID,
		)
		return err
	})
}

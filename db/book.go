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
	description=?, totalCopies=? , currentCopies=? where id = ?`

	bookidexist = `SELECT COUNT(*) FROM book WHERE book.id = ?`

	//GetTotalCopiesQuery  = `SELECT totalCopies FROM book where id=? `

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

// func (s *store) GetTotalCopies(ctx context.Context , bookId string) ( int){
// 	cnt:=0
// 	cnt, _ = s.db.Exec(GetTotalCopiesQuery, bookId)
// 	return cnt
// }

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
		var deleteTransaction []Request
		s.db.SelectContext(ctx, &deleteTransaction, listTransactionBookQueryByID, id)
		cnt := 0
		for _, j := range deleteTransaction {
			if j.Returndate == 0 {
				cnt++
			}
		}
		if cnt != 0 {
			return ErrBookExistTransaction
		}
		_, err := s.db.Exec(deleteTransactionBookQuery, id)
		if err != nil {
			return err
		}
		res, _ := s.db.Exec(deleteBookByIDQuery, id)
		if err != nil {
			return err
		}
		cnt1, err := res.RowsAffected()
		if err != nil {
			return err
		}
		if cnt1 == 0 {
			return ErrBookNotExist
		}
		return err
	})
}

func (s *store) UpdateBook(ctx context.Context, book *Book) (err error) {
	flag := 0

	s.db.GetContext(ctx, &flag, bookidexist, book.ID)

	if flag == 0 {
		return ErrIDNotExist
	} else {

		return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
			var tempBook Book
			s.db.GetContext(ctx, &tempBook, findBookByIDQuery, book.ID)
			if tempBook.TotalCopies > book.TotalCopies {
				return ErrLessThanPreviousTotal
			}
			totalcnt := 0
			currentcnt := 0
			s.db.GetContext(ctx, &totalcnt, GetTotalCopiesQuery, book.ID)
			s.db.GetContext(ctx, &currentcnt, GetCurrentCopiesQuery, book.ID)
			book.CurrentCopies = currentcnt + (book.TotalCopies - totalcnt)
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
}

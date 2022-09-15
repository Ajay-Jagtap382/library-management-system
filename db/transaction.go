package db

import (
	"context"
	"database/sql"
	"time"
	// "math/big"
)

const (
	createTransactionQuery = `INSERT INTO transactions (
		id ,
    	issuedate ,
    	duedate ,
		returndate,
    	book_id ,
    	user_id 	
	)
        VALUES(?,?,?,?,?,?)`

	listTransactionQuery       = `SELECT * FROM transactions`
	deleteTransactionByIDQuery = `DELETE FROM transactions WHERE id = ?`
	updateTransactionQuery     = `UPDATE transactions SET returndate=? WHERE book_id = ? AND user_id =? AND returndate= 0`
)

type Transaction struct {
	ID         string `db:"id"`
	Issuedate  int    `db:"issuedate"`
	Duedate    int    `db:"duedate"`
	Returndate int    `db:"returndate"`
	Book_id    string `db:"book_id"`
	User_id    string `db:"user_id"`
}

func (s *store) CreateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	now := time.Now().UTC().Unix()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createTransactionQuery,
			transaction.ID,
			now,
			transaction.Duedate,
			0,
			transaction.Book_id,
			transaction.User_id,
		)
		return err
	})
}

func (s *store) ListTransaction(ctx context.Context) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &transactions, listTransactionQuery)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrTransactionNotExist
	}
	return
}

func (s *store) DeleteTransactionByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteTransactionByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrTransactionNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateTransactionQuery,
			transaction.Returndate,
			transaction.Book_id,
			transaction.User_id,
		)
		return err
	})
}

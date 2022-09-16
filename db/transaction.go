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
	issueCopyQuery             = `UPDATE book SET currentCopies=currentCopies-1 WHERE id = ? AND currentCopies>0`
	returnCopyQuery            = `UPDATE book SET currentCopies=currentCopies+1 WHERE id = ?`
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
	transaction.Duedate = int(now) + 864000

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
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			issueCopyQuery,
			transaction.Book_id,
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

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateTransactionQuery,
			time.Now().UTC().Unix(),
			transaction.Book_id,
			transaction.User_id,
		)
		if err != nil {
			return err
		}

		_, err = s.db.Exec(
			returnCopyQuery,
			transaction.Book_id,
		)
		return err
	})
}

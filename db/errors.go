package db

import "errors"

var (
	ErrUserNotExist          = errors.New("User does not exist in db")
	ErrBookNotExist          = errors.New("Book does not exist in db")
	ErrTransactionNotExist   = errors.New("Transaction does not exist in db")
	ErrIDNotExist            = errors.New("ID does not exist in db")
	ErrAlreadyReturn         = errors.New("Book has been already returned")
	ErrTakenUser             = errors.New("User has not returned the books which are issued to him please first return it then we can delete the profile.")
	ErrLessThanPreviousTotal = errors.New("New TotalCopies of books should be greater than previous TotalCopies")
	ErrBookExistTransaction  = errors.New("Book is taken by user first collect it then we can delete book")
)

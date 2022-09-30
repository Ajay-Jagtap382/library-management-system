package book

import "errors"

var (
	errEmptyName             = errors.New("Book name must be present")
	errEmptyDescription      = errors.New("Description must be present")
	errZeroCopies            = errors.New("Copies cannot be zero while creation of book")
	errNoBooks               = errors.New("No book present")
	errNoBookId              = errors.New("Book is not present")
	ErrLessThanPreviousTotal = errors.New("New TotalCopies of books should be greater than previous TotalCopies")
	ErrBookExistTransaction  = errors.New("Book is taken by user first collect it then we can delete book")
)

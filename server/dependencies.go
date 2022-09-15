package server

import (
	"github.com/Ajay-Jagtap382/library-management-system/app"
	"github.com/Ajay-Jagtap382/library-management-system/book"
	"github.com/Ajay-Jagtap382/library-management-system/db"
	transaction "github.com/Ajay-Jagtap382/library-management-system/transaction"
	"github.com/Ajay-Jagtap382/library-management-system/users"
)

type dependencies struct {
	UserService        users.Service
	BookService        book.Service
	TransactionService transaction.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := users.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)
	transactionService := transaction.NewService(dbStore, logger)

	return dependencies{
		UserService:        userService,
		BookService:        bookService,
		TransactionService: transactionService,
	}, nil
}

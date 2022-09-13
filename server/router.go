package server

import (
	"net/http"

	"github.com/Ajay-Jagtap382/library-management-system/book"
	"github.com/Ajay-Jagtap382/library-management-system/users"

	"github.com/Ajay-Jagtap382/library-management-system/api"

	"github.com/gorilla/mux"
	//"github.com/Ajay-Jagtap382/library-management-system/config"
)

// const (
//  versionHeader = "Accept"
// )

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Users
	router.HandleFunc("/users", users.CreateUser(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", users.GetUser(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", users.GetUserByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users", users.UpdateUser(dep.UserService)).Methods(http.MethodPut)
	router.HandleFunc("/user/{userId}", users.DeleteUserByID(dep.UserService)).Methods(http.MethodDelete)

	//Books
	router.HandleFunc("/books", book.CreateBook(dep.BookService)).Methods(http.MethodPost)
	router.HandleFunc("/books", book.GetBook(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", book.GetBookByID(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books", book.UpdateBook(dep.BookService)).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", book.DeleteBookByID(dep.BookService)).Methods(http.MethodDelete)

	//Transactions
	// router.HandleFunc("/Transactions", book.CreateBook(dep.TransactionService)).Methods(http.MethodPost)
	// router.HandleFunc("/Transactions", book.GetBook(dep.TransactionService)).Methods(http.MethodGet)
	// router.HandleFunc("/Transactions", book.UpdateBook(dep.TransactionService)).Methods(http.MethodPut)
	// router.HandleFunc("/Transactions/{id}", book.DeleteBookByID(dep.TransactionService)).Methods(http.MethodDelete)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

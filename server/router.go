package server

import (
	// "encoding/json"
	// "errors"
	// "fmt"

	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Ajay-Jagtap382/library-management-system/book"
	transaction "github.com/Ajay-Jagtap382/library-management-system/transaction"
	"github.com/Ajay-Jagtap382/library-management-system/users"
	"github.com/golang-jwt/jwt"

	"github.com/Ajay-Jagtap382/library-management-system/api"

	"github.com/gorilla/mux"
)

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type TokenData struct {
	Id    string
	Email string
	Role  string
}

const (
	SUPERADMIN = iota
	ADMIN
	USER
)

var RoleMap = map[string]int{"superadmin": SUPERADMIN, "admin": ADMIN, "user": USER}

var jwtKey = []byte("jsd549$^&")

// var tokenRole = ""

// func Tokendatareturn() string {
// 	return tokenRole
// }

func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//TODO:
		//1. get token from reuqest header
		//2. decode token
		//3. check if user exist from token.ID
		//4. if role is allowed
		//5. call handler

		token := r.Header.Get("Authorization")

		isValid, TokenDatas, err := ValidateToken(token)
		fmt.Println(isValid)
		if err != nil {
			fmt.Println("error")
		}

		fmt.Println("Token Data : ", TokenDatas)

		if !isValid {
			api.Error(w, http.StatusBadRequest, api.Response{Message: "Token is not valid"})
			return
		}

		tokenRole := TokenDatas.Role
		if RoleMap[tokenRole] > role {
			api.Error(w, http.StatusBadRequest, api.Response{Message: "You don't have the access"})
			return
		}

		handler.ServeHTTP(w, r)
		return
	}
}

func ValidateToken(tokenString string) (isValid bool, tokenData TokenData, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	isValid = true

	tokenData = TokenData{
		Id:    claims.Id,
		Email: claims.Email,
		Role:  claims.Role,
	}
	return
}

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// login
	router.HandleFunc("/login", users.Login(dep.UserService)).Methods(http.MethodPost)

	// Users
	router.HandleFunc("/users", Authorize(users.CreateUser(dep.UserService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/users", Authorize(users.GetUser(dep.UserService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", Authorize(users.GetUserByID(dep.UserService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/users", Authorize(users.UpdateUser(dep.UserService), USER)).Methods(http.MethodPut)
	router.HandleFunc("/user/{userId}", Authorize(users.DeleteUserByID(dep.UserService), ADMIN)).Methods(http.MethodDelete)

	//Books
	router.HandleFunc("/books", Authorize(book.CreateBook(dep.BookService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/books", Authorize(book.GetBook(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/book/{id}", Authorize(book.GetBookByID(dep.BookService), USER)).Methods(http.MethodGet)
	router.HandleFunc("/books", Authorize(book.UpdateBook(dep.BookService), ADMIN)).Methods(http.MethodPut)
	router.HandleFunc("/book/{id}", Authorize(book.DeleteBookByID(dep.BookService), ADMIN)).Methods(http.MethodDelete)

	//Transactions
	router.HandleFunc("/Transactions", Authorize(transaction.CreateTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodPost)
	router.HandleFunc("/Transactions", Authorize(transaction.GetTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodGet)
	router.HandleFunc("/Transactions", Authorize(transaction.UpdateTransaction(dep.TransactionService), ADMIN)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

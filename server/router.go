package server

import (
	"net/http"

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

	// Category
	router.HandleFunc("/users", users.CreateUser(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", users.GetUser(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", users.GetUserByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/user/{userId}", users.UpdateUser(dep.UserService)).Methods("PUT")
	router.HandleFunc("/user/{userId}", users.DeleteUserByID(dep.UserService)).Methods(http.MethodDelete)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

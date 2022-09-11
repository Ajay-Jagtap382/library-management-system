package server

import (
	"net/http"

	//"github.com/Ajay-Jagtap382/Library-Management-System/server/users/handler"

	"github.com/Ajay-Jagtap382/Library-Management-System/api"

	"github.com/gorilla/mux"
	//"github.com/Ajay-Jagtap382/Library-Management-System/config"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())
	// TODO: add doc
	// v2 := fmt.Sprintf("application/vnd.%s.v2", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	// Category
	// router.HandleFunc("/users/", handler.CreteUser).Methods("POST")
	// router.HandleFunc("/users/", handler.GetUser).Methods("GET")
	// router.HandleFunc("/user/{userId}", handler.GetUserById).Methods("GET")
	// router.HandleFunc("/user/{userId}", handler.UpdateUser).Methods("PUT")
	// router.HandleFunc("/user/{userId}", handler.DeleteUser).Methods("DELETE")

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}

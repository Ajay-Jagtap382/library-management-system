package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Ajay-Jagtap382/library-management-system/api"
	"github.com/gorilla/mux"
)

// Transaction
func CreateTransaction(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c Request
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Create(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func unixToDate(uni int) string {
	unikey := strconv.Itoa(uni)
	i, _ := strconv.ParseInt(unikey, 10, 64)
	tm := time.Unix(i, 0)
	dateString := tm.Format("2006-01-02 15:04:05")
	return dateString
}

func GetTransaction(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.List(req.Context())
		if err == errNoTransaction {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var respo []Transactionresp
		for _, j := range resp.Transaction {

			var res Transactionresp
			res.ID = j.ID
			res.Issuedate = unixToDate(j.Issuedate)
			if j.Returndate != 0 {
				res.Returndate = unixToDate(j.Returndate)
			} else {
				res.Returndate = strconv.Itoa(j.Returndate)
			}
			res.Duedate = unixToDate(j.Duedate)
			res.Book_id = j.Book_id
			res.User_id = j.User_id
			respo = append(respo, res)
		}

		api.Success(rw, http.StatusOK, respo)
	})
}

func GetTransactionByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.ListByID(req.Context(), vars["id"])
		if err == errNoTransaction {
			api.Error(rw, http.StatusNotFound, api.Response{Message: errNoTransaction.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var respo []Transactionresp
		for _, j := range resp.Transaction {

			var res Transactionresp
			res.ID = j.ID
			res.Issuedate = unixToDate(j.Issuedate)
			if j.Returndate != 0 {
				res.Returndate = unixToDate(j.Returndate)
			} else {
				res.Returndate = strconv.Itoa(j.Returndate)
			}
			res.Duedate = unixToDate(j.Duedate)
			res.Book_id = j.Book_id
			res.User_id = j.User_id
			respo = append(respo, res)
		}

		api.Success(rw, http.StatusOK, respo)
	})
}

func GetBookStatus(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c RequestStatus
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		resp, err := service.BookStatus(req.Context(), c)
		if err == errNoTransaction {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}
		fmt.Println(resp)

		api.Success(rw, http.StatusOK, resp)
	})
}

func UpdateTransaction(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c Request
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})
}

func isBadRequest(err error) bool {
	return err == errEmptyID
}

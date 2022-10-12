package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Ajay-Jagtap382/library-management-system/api"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type TokenData struct {
	Id    string
	Email string
	Role  string
}

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("jsd549$^&")

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

func CreateBook(service Service) http.HandlerFunc {
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

func GetBook(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)
		resp, err := service.List(req.Context())
		if err == errNoBooks {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		var resposlice []Bookres

		for _, j := range resp.Books {
			reso, _ := service.BookStatus(req.Context(), j.ID, TokenDatas.Id)
			var respo Bookres
			respo.BookName = j.BookName
			respo.Description = j.Description
			respo.BookStatus = reso

			resposlice = append(resposlice, respo)
		}

		api.Success(rw, http.StatusOK, resposlice)
	})
}

func GetBookByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)
		vars := mux.Vars(req)

		resp, err := service.FindByID(req.Context(), vars["id"])

		if err == errNoBookId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		reso, _ := service.BookStatus(req.Context(), resp.Book.ID, TokenDatas.Id)

		var Bookrespo Bookres

		Bookrespo.BookName = resp.Book.BookName
		Bookrespo.Description = resp.Book.Description
		Bookrespo.BookStatus = reso

		api.Success(rw, http.StatusOK, Bookrespo)
	})
}

func DeleteBookByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.DeleteByID(req.Context(), vars["id"])
		if err == errNoBookId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}

func UpdateBook(service Service) http.HandlerFunc {
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
	return err == errEmptyName
}

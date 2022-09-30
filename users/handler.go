package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"net/http"

	"github.com/Ajay-Jagtap382/library-management-system/api"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

var cs UserService

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

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var j Authentication
		err := json.NewDecoder(req.Body).Decode(&j)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		jwtString, err1 := service.GenerateJWT(req.Context(), j.Email, j.Password)
		if err1 != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err1.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)

	})
}

func CreateUser(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c CreateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		err = c.Validate()
		if err != nil {
			//cs.logger.Errorw("Invalid request for user Create", "msg", err.Error(), "user", c)
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)

		if TokenDatas.Role == "admin" && c.Role == "admin" {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: errNotAccess.Error()})
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

func GetUser(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		resp, err := service.List(req.Context())
		if err == errNoUsers {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func GetUserByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)

		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)

		// if TokenDatas.Id != vars["id"] {
		// 	api.Error(rw, http.StatusBadRequest, api.Response{Message: errWrongUser.Error()})
		// 	return
		// }

		resp, err := service.FindByID(req.Context(), TokenDatas.Id)

		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func UpdatePassword(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)
		var c ChangePassword
		err := json.NewDecoder(req.Body).Decode(&c)

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		if len(c.NewPassword) < 6 {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: errMinimumLengthPassword.Error()})
			return
		}
		resp, err := service.List(req.Context())

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		flag := false

		for _, v := range resp.Users {
			if v.ID == TokenDatas.Id && v.Password == c.Password {
				flag = true
				err = service.UpdatePassword(req.Context(), c, TokenDatas)
				if isBadRequest(err) {
					api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
					return
				}

				if err != nil {
					api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
					return
				}

				api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
			}
		}
		if !flag {
			api.Success(rw, http.StatusOK, api.Response{Message: "Wrong ID or pasword"})
		}

	})
}

func DeleteUserByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)

		// fmt.Println(TokenDatas.Role)
		// fmt.Println(vars["userId"])
		role, err := service.FindByID(req.Context(), vars["userId"])
		if err != nil {
			api.Error(rw, http.StatusNotFound, api.Response{Message: errNoUsers.Error()})
			return
		}

		if RoleMap[TokenDatas.Role] >= RoleMap[role.User.Role] {
			api.Error(rw, http.StatusNotFound, api.Response{Message: errNotAccess.Error()})
			return
		}

		err = service.DeleteByID(req.Context(), vars["userId"])
		if err == errNoUserId {
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

func UpdateUser(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		_, TokenDatas, _ := ValidateToken(token)
		var c UpdateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Update(req.Context(), c, TokenDatas)
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
	return err == errEmptyFirstName || err == errEmptyID
}

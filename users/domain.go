package users

import (
	"net/mail"

	"github.com/Ajay-Jagtap382/library-management-system/db"
)

type updateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
}

type createRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
}

type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	Users []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.First_Name == "" {
		return errEmptyFirstName
	}
	if cr.Last_Name == "" {
		return errEmptyLastName
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if cr.Gender == "" {
		return errEmptyGender
	}
	if cr.Email == "" {
		return errEmptyEmail
	}
	if cr.Mobile_Num == "" {
		return errEmptyMobNo
	}
	if cr.Role == "" {
		return errEmptyRole
	}
	if cr.Role != "user" && cr.Role != "admin" && cr.Role != "superadmin" {
		return errRoleType
	}
	_, b := mail.ParseAddress(cr.Email)
	if b != nil {
		return errNotValidMail
	}
	if len(cr.Mobile_Num) < 10 || len(cr.Mobile_Num) > 10 {
		return errInvalidMobNo
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyFirstName
	}
	return
}

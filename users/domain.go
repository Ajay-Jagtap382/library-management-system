package users

import (
	"strings"
	"unicode"

	"github.com/Ajay-Jagtap382/library-management-system/db"
)

type UpdateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
}

type CreateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
	Role       string `json:"role"`
}

type FindByIDResponse struct {
	User db.User `json:"user"`
}

type ChangePassword struct {
	ID          string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type ListResponse struct {
	Users []db.User `json:"users"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.First_Name == "" {
		return errEmptyFirstName
	}
	if cr.Last_Name == "" {
		return errEmptyLastName
	}
	for _, j := range cr.First_Name {
		if !unicode.IsLetter(j) {
			return errInvalidFirstName
		}
	}
	for _, j := range cr.Last_Name {
		if !unicode.IsLetter(j) {
			return errInvalidLastName
		}
	}
	if cr.Password == "" {
		return errEmptyPassword
	}
	if len(cr.Password) < 6 {
		return errMinimumLengthPassword
	}
	if cr.Gender == "" || (strings.ToLower(cr.Gender) != "male" && strings.ToLower(cr.Gender) != "female") {
		return errValideGender
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
	if strings.ToLower(cr.Role) != "user" && strings.ToLower(cr.Role) != "admin" {
		return errRoleType
	}
	// _, b := mail.ParseAddress(cr.Email)
	// if b != nil {
	// 	return errNotValidMail
	// }
	validateEmail := cr.Email
	flag := false
	lastapperance := 0
	for i := 0; i < len(validateEmail); i++ {
		if validateEmail[i] == ' ' {
			return errNotValidMail
		}
		if validateEmail[i] == '@' {
			flag = true
			lastapperance = i
		}
	}
	if !flag {
		return errNotValidMail
	}
	flag = false
	for i := lastapperance; i < len(validateEmail); i++ {
		if validateEmail[i] == ' ' {
			return errNotValidMail
		}
		if validateEmail[i] == '.' {
			flag = true
		}
	}
	if !flag {
		return errNotValidMail
	}
	if len(cr.Mobile_Num) < 10 || len(cr.Mobile_Num) > 10 {
		return errInvalidMobNo
	}
	for _, j := range cr.Mobile_Num {
		if !unicode.IsNumber(j) {
			return errInvalidMobNo
		}
	}
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyFirstName
	}
	return
}

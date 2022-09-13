package users

import "github.com/Ajay-Jagtap382/library-management-system/db"

type updateRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
}

type createRequest struct {
	ID         string `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Mobile_Num string `json:"mobile_num"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Gender     string `json:"gender"`
}

type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	Users []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.First_Name == "" {
		return errEmptyName
	}
	return
}

func (ur updateRequest) Validate() (err error) {
	if ur.ID == "" {
		return errEmptyID
	}
	if ur.First_Name == "" {
		return errEmptyName
	}
	return
}

package users

import "github.com/Ajay-Jagtap382/library-management-system/db"

type Request struct {
	ID         string `json:"id"`
	Issuedate  int    `json:"issuedate"`
	Duedate    int    `json:"duedate"`
	Returndate int    `json:"returndate"`
	Book_id    string `json:"book_id"`
	User_id    string `json:"user_id"`
}

type ListResponse struct {
	Transaction []db.Transaction `json:"transaction"`
}

func (cr Request) Validate() (err error) {
	if cr.ID == "" {
		return errEmptyID
	}
	return
}

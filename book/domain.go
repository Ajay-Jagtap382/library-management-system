package book

import "github.com/Ajay-Jagtap382/library-management-system/db"

type Request struct {
	ID            string `json:"id"`
	BookName      string `json:"BookName"`
	Description   string `json:"description"`
	TotalCopies   int    `json:"totalCopies"`
	CurrentCopies int    `json:"currentCopies"`
}

type findByIDResponse struct {
	Book db.Book `json:"book"`
}

type listResponse struct {
	Books []db.Book `json:"books"`
}

func (cr Request) Validate() (err error) {
	if cr.BookName == "" {
		return errEmptyName
	}
	return
}

// func (ur Request) Validate() (err error) {
// 	if ur.ID == "" {
// 		return errEmptyID
// 	}
// 	if ur.bookName == "" {
// 		return errEmptyName
// 	}
// 	return
// }

package users

import (
	"context"

	"github.com/Ajay-Jagtap382/library-management-system/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	BookStatus(ctx context.Context, c RequestStatus) (response string, err error)
	Create(ctx context.Context, req Request) (err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req Request) (err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *transactionService) List(ctx context.Context) (response ListResponse, err error) {
	transaction, err := cs.store.ListTransaction(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No Transaction present", "err", err.Error())
		return response, errNoTransaction
	}
	if err != nil {
		cs.logger.Error("Error listing Transactions", "err", err.Error())
		return
	}
	response.Transaction = transaction
	return
}

func (cs *transactionService) BookStatus(ctx context.Context, c RequestStatus) (response string, err error) {
	response, err = cs.store.BookStatus(ctx, c.BookID, c.UserID)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No Transaction present", "err", err.Error())
		return response, errNoTransaction
	}
	if err != nil {
		cs.logger.Error("Error listing Transactions", "err", err.Error())
		return
	}
	return
}

func (cs *transactionService) Create(ctx context.Context, c Request) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for transaction Create", "msg", err.Error(), "user", c)
		return
	}
	res, _ := cs.store.BookStatus(ctx, c.Book_id, c.User_id)
	if res == "issued" {
		return errAlreadyTaken
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()

	err = cs.store.CreateTransaction(ctx, &db.Transaction{

		ID:         c.ID,
		Issuedate:  c.Issuedate,
		Duedate:    c.Duedate,
		Returndate: c.Returndate,
		Book_id:    c.Book_id,
		User_id:    c.User_id,
	})
	if err != nil {
		cs.logger.Error("Error creating transaction", "err", err.Error())
		return
	}
	return
}

func (cs *transactionService) Update(ctx context.Context, c Request) (err error) {
	if err != nil {
		cs.logger.Error("Invalid Request for transaction Update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateTransaction(ctx, &db.Transaction{
		Returndate: c.Returndate,
		Book_id:    c.Book_id,
		User_id:    c.User_id,
	})
	if err != nil {
		cs.logger.Error("Error updating transaction", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *transactionService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("user Not present", "err", err.Error(), "user_id", id)
		return errNoTransactionId
	}
	if err != nil {
		cs.logger.Error("Error deleting Transaction", "err", err.Error(), "user_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &transactionService{
		store:  s,
		logger: l,
	}
}

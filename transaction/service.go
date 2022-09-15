package users

import (
	"context"

	"github.com/Ajay-Jagtap382/library-management-system/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req Request) (err error)
	deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req Request) (err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *transactionService) list(ctx context.Context) (response listResponse, err error) {
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

func (cs *transactionService) create(ctx context.Context, c Request) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for transaction create", "msg", err.Error(), "user", c)
		return
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

func (cs *transactionService) update(ctx context.Context, c Request) (err error) {
	if err != nil {
		cs.logger.Error("Invalid Request for transaction update", "err", err.Error(), "user", c)
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

func (cs *transactionService) deleteByID(ctx context.Context, id string) (err error) {
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
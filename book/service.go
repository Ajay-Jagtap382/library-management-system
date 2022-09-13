package book

import (
	"context"

	"github.com/Ajay-Jagtap382/library-management-system/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req Request) (err error)
	findByID(ctx context.Context, id string) (response findByIDResponse, err error)
	deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req Request) (err error)
}

type bookService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *bookService) list(ctx context.Context) (response listResponse, err error) {
	books, err := cs.store.ListBooks(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No book present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing categories", "err", err.Error())
		return
	}
	response.Books = books
	return
}

func (cs *bookService) create(ctx context.Context, c Request) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for user create", "msg", err.Error(), "user", c)
		return
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()

	err = cs.store.CreateBook(ctx, &db.Book{
		ID:            c.ID,
		BookName:      c.BookName,
		Description:   c.Description,
		TotalCopies:   c.TotalCopies,
		CurrentCopies: c.CurrentCopies,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func (cs *bookService) update(ctx context.Context, c Request) (err error) {
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateBook(ctx, &db.Book{
		ID:            c.ID,
		BookName:      c.BookName,
		Description:   c.Description,
		TotalCopies:   c.TotalCopies,
		CurrentCopies: c.CurrentCopies,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *bookService) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
	book, err := cs.store.FindBookByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "user_id", id)
		return
	}

	response.Book = book
	return
}

func (cs *bookService) deleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteBookByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("user Not present", "err", err.Error(), "user_id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "user_id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &bookService{
		store:  s,
		logger: l,
	}
}

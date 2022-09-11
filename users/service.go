package users

// import (
// 	"context"

// 	"github.com/Ajay-Jagtap382/Library-Management-System/db"

// 	"go.uber.org/zap"
// )

// type Service interface {
// 	list(ctx context.Context) (response listResponse, err error)
// 	create(ctx context.Context, req createRequest) (err error)
// 	findByID(ctx context.Context, id string) (response findByIDResponse, err error)
// 	deleteByID(ctx context.Context, id string) (err error)
// 	update(ctx context.Context, req updateRequest) (err error)
// }

// // type categoryService struct {
// // 	store  db.Storer
// // 	logger *zap.SugaredLogger
// // }

// type enduser struct {
// 	store  db.Storer
// 	logger *zap.SugaredLogger
// 	id     int    `json:"id"`
// 	Email  string `json:"email"`
// 	name   string `json:"name"`
// 	phone  string `json:"phoneNum"`
// }

// func (cs *enduser) GetUser(ctx context.Context) (response listResponse, err error) {
// 	users, err := cs.store.ListCategories(ctx)
// 	if err == db.ErrCategoryNotExist {
// 		cs.logger.Error("No users present", "err", err.Error())
// 		return response, err
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error listing users", "err", err.Error())
// 		return
// 	}

// 	response.Categories = users
// 	return
// }

// func (cs *enduser) create(ctx context.Context, c createRequest) (err error) {
// 	err = c.Validate()
// 	if err != nil {
// 		cs.logger.Errorw("Invalid request for category create", "msg", err.Error(), "category", c)
// 		return
// 	}

// 	err = cs.store.CreateCategory(ctx, &db.Category{
// 		Name: c.Name,
// 	})
// 	if err != nil {
// 		cs.logger.Error("Error creating category", "err", err.Error())
// 		return
// 	}
// 	return
// }

// func (cs *enduser) update(ctx context.Context, c updateRequest) (err error) {
// 	err = c.Validate()
// 	if err != nil {
// 		cs.logger.Error("Invalid Request for category update", "err", err.Error(), "category", c)
// 		return
// 	}

// 	err = cs.store.UpdateCategory(ctx, &db.Category{
// 		ID:   c.ID,
// 		Name: c.Name,
// 	})
// 	if err != nil {
// 		cs.logger.Error("Error updating category", "err", err.Error(), "category", c)
// 		return
// 	}

// 	return
// }

// func (cs *enduser) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
// 	category, err := cs.store.FindCategoryByID(ctx, id)
// 	if err == db.ErrCategoryNotExist {
// 		cs.logger.Error("No category present", "err", err.Error())
// 		return response, errNoCategoryId
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error finding category", "err", err.Error(), "category_id", id)
// 		return
// 	}

// 	response.Category = category
// 	return
// }

// func (cs *enduser) deleteByID(ctx context.Context, id string) (err error) {
// 	err = cs.store.DeleteCategoryByID(ctx, id)
// 	if err == db.ErrCategoryNotExist {
// 		cs.logger.Error("Category Not present", "err", err.Error(), "category_id", id)
// 		return errNoCategoryId
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error deleting category", "err", err.Error(), "category_id", id)
// 		return
// 	}

// 	return
// }

// func NewService(s db.Storer, l *zap.SugaredLogger) Service {
// 	return &categoryService{
// 		store:  s,
// 		logger: l,
// 	}
// }

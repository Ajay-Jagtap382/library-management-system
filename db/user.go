package db

// import (
// 	"context"
// 	"database/sql"
// 	// "time"
// )

// const (
// 	createCategoryQuery = `INSERT INTO categories (
//         name, created_at, updated_at)
//         VALUES($1, $2, $3)
//         `
// 	listCategoriesQuery     = `SELECT * FROM categories`
// 	findCategoryByIDQuery   = `SELECT * FROM categories WHERE id = $1`
// 	deleteCategoryByIDQuery = `DELETE FROM categories WHERE id = $1`
// 	updateCategoryQuery     = `UPDATE categories SET name = $1, updated_at = $2 where id = $3`
// )

// // type Category struct {
// // 	ID        string    `db:"id"`
// // 	Name      string    `db:"name"`
// // 	CreatedAt time.Time `db:"created_at"`
// // 	UpdatedAt time.Time `db:"updated_at"`
// // }
// type Enduser struct {
// 	id    int    `json:"id"`
// 	Email string `json:"email"`
// 	name  string `json:"name"`
// 	phone string `json:"phoneNum"`
// }

// // func (s *store) CreateCategory(ctx context.Context, category *Category) (err error) {
// // 	now := time.Now()

// // 	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
// // 		_, err = s.db.Exec(
// // 			createCategoryQuery,
// // 			category.Name,
// // 			now,
// // 			now,
// // 		)
// // 		return err
// // 	})
// // }

// func (s *store) ListCategories(ctx context.Context) (categories []Enduser, err error) {
// 	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
// 		return s.db.SelectContext(ctx, &categories, listCategoriesQuery)
// 	})
// 	if err == sql.ErrNoRows {
// 		return categories, ErrCategoryNotExist
// 	}
// 	return
// }

// // func (s *store) FindCategoryByID(ctx context.Context, id string) (category Category, err error) {
// // 	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
// // 		return s.db.GetContext(ctx, &category, findCategoryByIDQuery, id)
// // 	})
// // 	if err == sql.ErrNoRows {
// // 		return category, ErrCategoryNotExist
// // 	}
// // 	return
// // }

// // func (s *store) DeleteCategoryByID(ctx context.Context, id string) (err error) {
// // 	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
// // 		res, err := s.db.Exec(deleteCategoryByIDQuery, id)
// // 		cnt, err := res.RowsAffected()
// // 		if cnt == 0 {
// // 			return ErrCategoryNotExist
// // 		}
// // 		if err != nil {
// // 			return err
// // 		}
// // 		return err
// // 	})
// // }

// // func (s *store) UpdateCategory(ctx context.Context, category *Category) (err error) {
// // 	now := time.Now()

// // 	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
// // 		_, err = s.db.Exec(
// // 			updateCategoryQuery,
// // 			category.Name,
// // 			now,
// // 			category.ID,
// // 		)
// // 		return err
// // 	})
// // }

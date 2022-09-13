package db

import (
	"context"
	"database/sql"
)

const (
	createUserQuery = `INSERT INTO user (
		id,
		first_name,
		last_name,
		mobile_num ,
		email,
		password,
		gender)
        VALUES(?,?,?,?,?,?,?)`

	listUserQuery       = `SELECT * FROM user`
	findUserByIDQuery   = `SELECT * FROM user WHERE id = ?`
	deleteUserByIDQuery = `DELETE FROM user WHERE id = ?`
	updateUserQuery     = `UPDATE user SET first_name = ?, last_name = ? , 
	mobile_num = ? ,gender=? ,password=?  WHERE id = ?`
)

type User struct {
	ID         string `db:"id"`
	First_Name string `db:"first_name"`
	Last_Name  string `db:"last_name"`
	Mobile_Num string `db:"mobile_num"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Gender     string `db:"gender"`
}

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.ID,
			user.First_Name,
			user.Last_Name,
			user.Mobile_Num,
			user.Email,
			user.Password,
			user.Gender,
		)
		return err
	})
}

func (s *store) ListUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, listUserQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}
	return
}

func (s *store) FindUserByID(ctx context.Context, id string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (s *store) DeleteUserByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteUserByIDQuery, id)
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrUserNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.First_Name,
			user.Last_Name,
			user.Mobile_Num,
			user.Gender,
			user.Password,
			user.ID,
		)
		return err
	})
}

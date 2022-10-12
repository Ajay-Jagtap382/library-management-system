package users

import (
	"context"
	// "crypto/aes"
	// "crypto/cipher"
	// "encoding/base64"
	"fmt"
	"time"

	"github.com/Ajay-Jagtap382/library-management-system/db"
	// "github.com/Ajay-Jagtap382/library-management-system/server"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	Create(ctx context.Context, req CreateRequest) (err error)
	FindByID(ctx context.Context, id string) (response FindByIDResponse, err error)
	GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req UpdateRequest, TokenDatas TokenData) (err error)
	UpdatePassword(ctx context.Context, req ChangePassword, TokenDatas TokenData) (err error)
}

type UserService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("jsd549$^&")

func (cs *UserService) GenerateJWT(ctx context.Context, Email string, Password string) (tokenString string, err error) {

	// var cs *userService
	user, err := cs.store.FindUserByEmail(ctx, Email)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return "", errNoUserId
	}
	// if err != nil {
	// 	cs.logger.Error("Error finding user", "err", err.Error(), "email", Email)
	// 	return "", err
	// }
	dbPassKey := []byte(user.Password)
	getpasskey := []byte(Password)
	if bcrypt.CompareHashAndPassword(dbPassKey, getpasskey) != nil {
		// cs.logger.Error("Wrong Password", "err", err.Error())
		return "", errWrongPassword
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (cs *UserService) List(ctx context.Context) (response ListResponse, err error) {
	users, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No category present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing categories", "err", err.Error())
		return
	}
	response.Users = users
	return
}

// var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// // This should be in an env file in production
// const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

// func Encode(b []byte) string {
// 	return base64.StdEncoding.EncodeToString(b)
// }

// // Encrypt method is to encrypt or hide any classified text
// func Encrypt(text, MySecret string) (string, error) {
// 	block, err := aes.NewCipher([]byte(MySecret))
// 	if err != nil {
// 		return "", err
// 	}
// 	plainText := []byte(text)
// 	cfb := cipher.NewCFBEncrypter(block, bytes)
// 	cipherText := make([]byte, len(plainText))
// 	cfb.XORKeyStream(cipherText, plainText)
// 	return Encode(cipherText), nil
// }

func (cs *UserService) Create(ctx context.Context, c CreateRequest) (err error) {
	// err = c.Validate()
	// if err != nil {
	// 	cs.logger.Errorw("Invalid request for user Create", "msg", err.Error(), "user", c)
	// 	return
	// }

	// if server.Tokendatareturn() == "admin" && c.Role == "superadmin" {
	// 	return errCreateSuperadmin
	// }

	uuidgen := uuid.New()
	c.ID = uuidgen.String()

	// PasswordEnrc, err := Encrypt(c.Password, MySecret)

	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}

	passkey := []byte(c.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passkey, bcrypt.DefaultCost)

	err = cs.store.CreateUser(ctx, &db.User{
		ID:         c.ID,
		First_Name: c.First_Name,
		Last_Name:  c.Last_Name,
		Mobile_Num: c.Mobile_Num,
		Email:      c.Email,
		Password:   string(hashedPassword),
		Gender:     c.Gender,
		Role:       c.Role,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func (cs *UserService) Update(ctx context.Context, c UpdateRequest, TokenDatas TokenData) (err error) {
	if err != nil {
		cs.logger.Error("Invalid Request for user Update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		First_Name: c.First_Name,
		Last_Name:  c.Last_Name,
		ID:         TokenDatas.Id,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *UserService) UpdatePassword(ctx context.Context, c ChangePassword, TokenDatas TokenData) (err error) {
	// err = c.Validate()
	// if err != nil {
	//  cs.logger.Error("Invalid Request for Password update", "err", err.Error(), "user", c)
	//  return
	// }

	passkey := []byte(c.NewPassword)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passkey, bcrypt.DefaultCost)
	err = cs.store.UpdatePassword(ctx, &db.User{
		ID:       TokenDatas.Id,
		Password: string(hashedPassword),
	})
	if err != nil {
		cs.logger.Error("Error updating Password", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *UserService) FindByID(ctx context.Context, id string) (response FindByIDResponse, err error) {
	user, err := cs.store.FindUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "user_id", id)
		return
	}

	response.User = user
	return
}

func (cs *UserService) DeleteByID(ctx context.Context, id string) (err error) {
	// var deleteUser trans.Request
	// deleteUser,err1 := trans.ListByID(ctx , id)
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("user Not present", "err", err.Error(), "user_id", id)
		return errNoUserId
	}
	if err == db.ErrTakenUser {
		//cs.logger.Error("user Not present", "err", err.Error(), "user_id", id)
		return errTakenUser
	}
	if err != nil {
		// cs.logger.Error("Error deleting user", "err", err.Error(), "user_id", id)
		return err
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &UserService{
		store:  s,
		logger: l,
	}
}

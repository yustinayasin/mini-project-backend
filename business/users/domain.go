package users

import (
	"context"
	"time"

	"gorm.io/gorm"
)

//stuct untuk keluar masuk usecase
type User struct {
	Id          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Street      string
	Address     string
	PostalCode  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm"index"`
	Token       string
}

//ini interface antara controller dan usecase
type UserUseCaseInterface interface {
	//perlu konteks biar ada timeout
	Login(user User, ctx context.Context) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	SignUp(user User, ctx context.Context) (User, error)
	GetUserDetail(id int, ctx context.Context) (User, error)
	EditUser(user User, id int, ctx context.Context) (User, error)
	DeleteUser(id int, ctx context.Context) (User, error)
}

//ini interface antara usecase dan repo
type UserRepoInterface interface {
	Login(user User, ctx context.Context) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	SignUp(user User, ctx context.Context) (User, error)
	GetUserDetail(id int, ctx context.Context) (User, error)
	EditUser(user User, id int, ctx context.Context) (User, error)
	DeleteUser(id int, ctx context.Context) (User, error)
}

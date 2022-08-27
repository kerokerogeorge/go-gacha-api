package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	CreateUser(name string, token string) (string, error)
	GetUser(token string) (*model.User, error)
	UpdateUser(user *model.User, name string) (*model.User, error)
	DeleteUser(user *model.User) error
}

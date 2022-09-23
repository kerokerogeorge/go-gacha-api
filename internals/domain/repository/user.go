package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type UserRepository interface {
	CreateUser(name string, token string, address string) (string, error)
	GetUser(token string) (*model.User, error)
	GetUsers() ([]*model.User, error)
	UpdateUser(user *model.User, name string) (*model.User, error)
	DeleteUser(user *model.User) error
}

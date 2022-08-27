package repository

type UserRepository interface {
	CreateUser(name string, token string) (string, error)
	GetUser(token string) (string, error)
}

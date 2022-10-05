package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type UserCharcacterRepository interface {
	GetResults(userId string) ([]*model.Result, error)
	CreateUserCharacter(*model.UserCharacter) error
	GetUserCharacters(id string, queryType string) ([]*model.UserCharacter, error)
	DeleteUserCharacter(userCharacter *model.UserCharacter) error
}

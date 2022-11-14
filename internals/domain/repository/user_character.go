package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type UserCharcacterRepository interface {
	GetResults(userId string) ([]*model.Result, error)
	CreateUserCharacter(*model.UserCharacter) (*model.UserCharacter, error)
	GetUserCharacters(id string, queryType string) ([]*model.UserCharacter, error)
	DeleteUserCharacter(userCharacter *model.UserCharacter) error
	GetOne(userCharacterID uint) (*model.UserCharacter, error)
	UpdateUsercharacter(userCharacter *model.UserCharacter, status model.CharacterStatus) (*model.UserCharacter, error)
}

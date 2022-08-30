package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

// import (
// "context"
// "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

// "github.com/gin-gonic/gin"
// )

type CharacterRepository interface {
	GetCharacters() ([]*model.Character, error)
	CreateCharacter(character *model.Character) (*model.Character, error)
	GetCharacter(characterId int) (*model.Character, error)
	CreateUserCharacter(*model.Result) error
}

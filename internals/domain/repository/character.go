package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type CharacterRepository interface {
	GetCharacters() ([]*model.Character, error)
	CreateCharacter(character *model.Character) (*model.Character, error)
	GetCharacter(characterId int) (*model.Character, error)
	DeleteCharacter(character *model.Character) error
}

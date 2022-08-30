package repository

import (
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type CharacterEmmitionRateRepository interface {
	SetEmmitionRate(characterWithEmmitionRate *model.CharacterEmmitionRate) error
	GetCharacterWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error)
}

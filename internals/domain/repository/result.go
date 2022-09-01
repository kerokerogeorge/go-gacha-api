package repository

import "github.com/kerokerogeorge/go-gacha-api/internals/domain/model"

type ResultRepository interface {
	GetResults(userId string) ([]*model.UserCharacter, error)
}

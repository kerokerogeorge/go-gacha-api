package usecase

import (
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type CharacterUsecase interface {
	GetCharacters() ([]*model.Character, error)
}

type characterUsecase struct {
	characterRepo repository.CharacterRepository
}

func NewCharacterUsecase(cr repository.CharacterRepository) CharacterUsecase {
	return &characterUsecase{
		characterRepo: cr,
	}
}

func (cu *characterUsecase) GetCharacters() ([]*model.Character, error) {
	return cu.characterRepo.GetCharacters()
}

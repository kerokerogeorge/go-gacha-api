package usecase

import (
	"log"

	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type CharacterUsecase interface {
	GetCharacters() ([]*model.Character, error)
	Create(character *model.Character) (*model.Character, error)
	GetCharactersWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error)
}

type characterUsecase struct {
	characterRepo             repository.CharacterRepository
	characterEmmitionRateRepo repository.CharacterEmmitionRateRepository
}

func NewCharacterUsecase(cr repository.CharacterRepository, cerr repository.CharacterEmmitionRateRepository) CharacterUsecase {
	return &characterUsecase{
		characterRepo:             cr,
		characterEmmitionRateRepo: cerr,
	}
}

func (cu *characterUsecase) GetCharacters() ([]*model.Character, error) {
	return cu.characterRepo.GetCharacters()
}

func (cu *characterUsecase) Create(character *model.Character) (*model.Character, error) {
	return cu.characterRepo.CreateCharacter(character)
}

func (cu *characterUsecase) GetCharactersWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error) {
	log.Println("gachaId", gachaId)
	return cu.characterEmmitionRateRepo.GetCharacterWithEmmitionRate(gachaId)
}

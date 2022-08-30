package usecase

import (
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type GachaUsecase interface {
	Create(gacha *model.Gacha) (*model.Gacha, error)
	List() ([]*model.Gacha, error)
	Get(gachaId string) (*model.Gacha, error)
}

type gachaUsecase struct {
	gachaRepo                 repository.GachaRepository
	characterRepo             repository.CharacterRepository
	characterEmmitionRateRepo repository.CharacterEmmitionRateRepository
}

func NewGachaUsecase(gr repository.GachaRepository, cr repository.CharacterRepository, cerr repository.CharacterEmmitionRateRepository) GachaUsecase {
	return &gachaUsecase{
		gachaRepo:                 gr,
		characterRepo:             cr,
		characterEmmitionRateRepo: cerr,
	}
}

func (gu *gachaUsecase) Create(gacha *model.Gacha) (*model.Gacha, error) {
	characters, err := gu.characterRepo.GetCharacters()
	if err != nil {
		return nil, err
	}

	if len(characters) == 0 {
		return nil, nil
	}

	// 排出率をキャラクターごとに出す
	for _, character := range characters {
		characterWithEmmitionRate, err := model.NewCharacterEmmitionRate(gacha.ID, character.ID)
		if err != nil {
			return nil, err
		}

		err = gu.characterEmmitionRateRepo.SetEmmitionRate(characterWithEmmitionRate)
		if err != nil {
			return nil, err
		}
	}

	return gu.gachaRepo.CreateGacha(gacha)
}

func (gu *gachaUsecase) List() ([]*model.Gacha, error) {
	return gu.gachaRepo.List()
}

func (gu *gachaUsecase) Get(gachaId string) (*model.Gacha, error) {
	return gu.gachaRepo.GetOne(gachaId)
}

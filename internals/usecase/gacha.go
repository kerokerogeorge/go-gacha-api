package usecase

import (
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/repository"
)

type GachaUsecase interface {
	Create(gacha *model.Gacha) (*model.Gacha, error)
	List() ([]*model.Gacha, error)
}

type gachaUsecase struct {
	gachaRepo repository.GachaRepository
}

func NewGachaUsecase(gr repository.GachaRepository) GachaUsecase {
	return &gachaUsecase{
		gachaRepo: gr,
	}
}

func (gu *gachaUsecase) Create(gacha *model.Gacha) (*model.Gacha, error) {
	return gu.gachaRepo.CreateGacha(gacha)
}

func (gu *gachaUsecase) List() ([]*model.Gacha, error) {
	return gu.gachaRepo.List()
}

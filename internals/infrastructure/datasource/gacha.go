package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type Gacha struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type gachaRepository struct {
	db *gorm.DB
}

func NewGachaRepository(database *gorm.DB) *gachaRepository {
	db := database
	return &gachaRepository{
		db: db,
	}
}

func (gr *gachaRepository) CreateGacha(gacha *model.Gacha) (*model.Gacha, error) {
	err := gr.db.Table("gachas").Create(gacha).Error
	if err != nil {
		return nil, err
	}
	return gacha, nil
}

func (gr *gachaRepository) List() ([]*model.Gacha, error) {
	var gachas []*model.Gacha
	err := gr.db.Find(&gachas).Error
	if err != nil {
		return nil, err
	}
	return gachas, nil
}

func (gr *gachaRepository) GetOne(gachaId string) (*model.Gacha, error) {
	var gacha Gacha
	err := gr.db.Table("gachas").Where("id = ?", gachaId).First(&gacha).Error
	if err != nil {
		return nil, err
	}
	return gr.ToGachaModel(gacha), nil
}

func (gr *gachaRepository) DeleteGacha(gacha *model.Gacha) error {
	err := gr.db.Delete(&gacha).Error
	if err != nil {
		return err
	}
	return nil
}

func (gr *gachaRepository) ToGachaModel(gacha Gacha) *model.Gacha {
	return &model.Gacha{
		ID:        gacha.ID,
		CreatedAt: gacha.CreatedAt,
		UpdatedAt: gacha.UpdatedAt,
	}
}

func (gr *gachaRepository) GetGachaCharacters(gachaId string) ([]*model.CharacterEmmitionRate, error) {
	var gachaCharacters []*model.CharacterEmmitionRate
	err := gr.db.Table("character_emmition_rates").Where("gacha_id = ?", gachaId).Find(&gachaCharacters).Error
	if err != nil {
		return nil, err
	}
	return gachaCharacters, nil
}

func (gr *gachaRepository) DeleteGachaCharacter(gachaCharacter *model.CharacterEmmitionRate) error {
	err := gr.db.Delete(&gachaCharacter).Error
	if err != nil {
		return err
	}
	return nil
}

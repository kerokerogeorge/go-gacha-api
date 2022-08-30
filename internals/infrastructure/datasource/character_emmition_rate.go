package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type CharacterEmmitionRate struct {
	ID           string    `json:"id"`
	GachaID      string    `json:"gachaId"`
	CharacterID  string    `json:"characterId"`
	EmissionRate int       `json:"emissionRate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type CharacterWithEmmitionRate struct {
	CharacterID  string `json:"characterId"`
	Name         string `json:"name"`
	EmissionRate string `json:"emissionRate"`
}

type characterEmmitionRateRepository struct {
	db *gorm.DB
}

func NewCharacterEmmitionRateRepository(database *gorm.DB) *characterEmmitionRateRepository {
	db := database
	return &characterEmmitionRateRepository{
		db: db,
	}
}

func (cerr *characterEmmitionRateRepository) SetEmmitionRate(characterWithEmmitionRate *model.CharacterEmmitionRate) error {
	err := cerr.db.Table("character_emmition_rates").Create(characterWithEmmitionRate).Error
	if err != nil {
		return err
	}
	return nil
}

func (cerr *characterEmmitionRateRepository) GetCharacterWithEmmitionRate(gachaId string) ([]*model.CharacterWithEmmitionRate, error) {
	var characterEmmitionRate []*model.CharacterWithEmmitionRate
	err := cerr.db.Table("gachas").Select("character_emmition_rates.character_id, characters.name, character_emmition_rates.emission_rate").
		Joins("INNER JOIN character_emmition_rates ON character_emmition_rates.gacha_id = ?", gachaId).
		Joins("INNER JOIN characters ON character_emmition_rates.character_id = characters.id").
		Where("gachas.id = ?", gachaId).
		Scan(&characterEmmitionRate).Error
	if err != nil {
		return nil, err
	}
	return characterEmmitionRate, nil
}

// func (cerr *characterEmmitionRateRepository) ToCharacterWithEmmitionRateModel(characterEmmitionRate []*CharacterWithEmmitionRate) *model.CharacterWithEmmitionRate {
// 	return &model.CharacterWithEmmitionRate{
// 		CharacterID:   characterEmmitionRate.CharacterID,
// 		CharacterName: characterEmmitionRate.Name,
// 		EmissionRate:  characterEmmitionRate.EmissionRate,
// 	}
// }

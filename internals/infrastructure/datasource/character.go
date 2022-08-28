package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type Character struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type characterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(database *gorm.DB) *characterRepository {
	db := database
	return &characterRepository{
		db: db,
	}
}

func (cr *characterRepository) GetCharacters() ([]*model.Character, error) {
	var characters []*model.Character
	err := cr.db.Find(&characters).Error
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func (cr *characterRepository) CreateCharacter(character *model.Character) (*model.Character, error) {
	err := cr.db.Table("characters").Create(character).Error
	if err != nil {
		return nil, err
	}
	return character, nil
}

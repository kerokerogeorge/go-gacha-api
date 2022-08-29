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

package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type Result struct {
	ID          string    `json:"id"`
	UserId      string    `json:"userId"`
	CharacterId string    `json:"characterId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type UserCharacter struct {
	ID          string `json:"usercharacterId"`
	CharacterId string `json:"characterId"`
	Name        string `json:"name"`
}

type resultRepository struct {
	db *gorm.DB
}

func NewResultRepository(database *gorm.DB) *resultRepository {
	db := database
	return &resultRepository{
		db: db,
	}
}

func (rr *resultRepository) GetResults(userId string) ([]*model.UserCharacter, error) {
	var results []*model.UserCharacter
	err := rr.db.Table("users").Select("user_characters.id, user_characters.character_id, characters.name").
		Joins("INNER JOIN user_characters ON user_characters.user_id = ?", userId).
		Joins("INNER JOIN characters ON user_characters.character_id = characters.id").
		Where("users.id = ?", userId).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

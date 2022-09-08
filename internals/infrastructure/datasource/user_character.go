package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type UserCharacter struct {
	ID          string    `json:"id"`
	UserId      string    `json:"userId"`
	CharacterId string    `json:"characterId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Result struct {
	ID          string `json:"usercharacterId"`
	CharacterId string `json:"characterId"`
	Name        string `json:"name"`
}

type userCharcacterRepository struct {
	db *gorm.DB
}

func NewResultRepository(database *gorm.DB) *userCharcacterRepository {
	db := database
	return &userCharcacterRepository{
		db: db,
	}
}

func (rr *userCharcacterRepository) GetResults(userId string) ([]*model.Result, error) {
	var results []*model.Result
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

func (rr *userCharcacterRepository) GetUserCharacters(id string, queryType string) ([]*model.UserCharacter, error) {
	var results []*model.UserCharacter
	table := rr.db.Table("user_characters")

	if queryType == "CHARACTER" {
		table = table.Where("character_id = ?", id)
	} else {
		table = table.Where("user_id = ?", id)
	}

	err := table.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (rr *userCharcacterRepository) DeleteUserCharacter(userCharacter *model.UserCharacter) error {
	err := rr.db.Delete(&userCharacter).Error
	if err != nil {
		return err
	}
	return nil
}
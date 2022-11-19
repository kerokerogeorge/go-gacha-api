package datasource

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kerokerogeorge/go-gacha-api/internals/domain/model"
)

type UserCharacter struct {
	ID           uint                  `json:"id" gorm:"primary_key"`
	UserId       string                `json:"userId"`
	CharacterId  string                `json:"characterId"`
	ImgUrl       string                `json:"imgUrl"`
	EmissionRate float64               `json:"emissionRate"`
	Status       model.CharacterStatus `json:"status"`
	CreatedAt    time.Time             `json:"createdAt"`
	UpdatedAt    time.Time             `json:"updatedAt"`
}

type userCharcacterRepository struct {
	db *gorm.DB
}

func NewUserCharacterRepository(database *gorm.DB) *userCharcacterRepository {
	db := database
	return &userCharcacterRepository{
		db: db,
	}
}

func (ucr *userCharcacterRepository) CreateUserCharacter(userCharacter *model.UserCharacter) (*model.UserCharacter, error) {
	err := ucr.db.Table("user_characters").Create(&userCharacter).Error
	if err != nil {
		return nil, err
	}
	return userCharacter, nil
}

func (ucr *userCharcacterRepository) GetResults(userId string) ([]*model.Result, error) {
	var results []*model.Result
	err := ucr.db.Table("users").Select("user_characters.id, user_characters.character_id, characters.name, user_characters.img_url, user_characters.emission_rate, user_characters.status").
		Joins("INNER JOIN user_characters ON user_characters.user_id = ?", userId).
		Joins("INNER JOIN characters ON user_characters.character_id = characters.id").
		Where("users.id = ?", userId).
		Where("user_characters.status = ?", model.CharacterStatusSuccess).
		Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ucr *userCharcacterRepository) GetUserCharacters(id string, queryType string) ([]*model.UserCharacter, error) {
	var results []*model.UserCharacter
	table := ucr.db.Table("user_characters")

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

func (ucr *userCharcacterRepository) DeleteUserCharacter(userCharacter *model.UserCharacter) error {
	err := ucr.db.Delete(&userCharacter).Error
	if err != nil {
		return err
	}
	return nil
}

func (ucr *userCharcacterRepository) GetOne(userCharacterID uint) (*model.UserCharacter, error) {
	var userCharacter UserCharacter
	err := ucr.db.Table("user_characters").Where("id = ?", userCharacterID).First(&userCharacter).Error
	if err != nil {
		return nil, err
	}
	return ucr.ToUserCharacterModel(userCharacter), nil
}

func (ucr *userCharcacterRepository) UpdateUsercharacter(userCharacter *model.UserCharacter, status model.CharacterStatus) (*model.UserCharacter, error) {
	database := ucr.db.Model(&userCharacter).Update("status", status)
	if database.Error != nil {
		return nil, database.Error
	}
	return userCharacter, nil
}

func (ucr *userCharcacterRepository) GetHistory(userId string) ([]*model.Result, error) {
	var results []*model.Result
	err := ucr.db.Table("users").Select("user_characters.id, user_characters.character_id, characters.name, user_characters.img_url, user_characters.emission_rate, user_characters.status").
		Joins("INNER JOIN user_characters ON user_characters.user_id = ?", userId).
		Joins("INNER JOIN characters ON user_characters.character_id = characters.id").
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ucr *userCharcacterRepository) ToUserCharacterModel(userCharacter UserCharacter) *model.UserCharacter {
	return &model.UserCharacter{
		ID:           userCharacter.ID,
		UserId:       userCharacter.UserId,
		CharacterId:  userCharacter.CharacterId,
		ImgUrl:       userCharacter.ImgUrl,
		EmissionRate: userCharacter.EmissionRate,
		Status:       userCharacter.Status,
		CreatedAt:    userCharacter.CreatedAt,
		UpdatedAt:    userCharacter.UpdatedAt,
	}
}

package model

import (
	"time"

	"github.com/Songmu/flextime"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserCharacter struct {
	ID          string    `json:"id" gorm:"primary_key"`
	UserId      string    `json:"userId"`
	CharacterId string    `json:"characterId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Result struct {
	ID          string `json:"userCharacterId"`
	CharacterId string `json:"characterId"`
	Name        string `json:"name"`
}

func NewUserCharacter(userId string, characterId string) (*UserCharacter, error) {
	now := flextime.Now()
	return &UserCharacter{
		UserId:      userId,
		CharacterId: characterId,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

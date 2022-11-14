package model

import (
	"time"

	"github.com/Songmu/flextime"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserCharacter struct {
	ID           string          `json:"id" gorm:"primary_key"`
	UserId       string          `json:"userId"`
	CharacterId  string          `json:"characterId"`
	ImgUrl       string          `json:"imgUrl"`
	EmissionRate float64         `json:"emissionRate"`
	Status       CharacterStatus `json:"status"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}

type CharacterStatus string

const (
	CharacterStatusPending CharacterStatus = "pending"
	CharacterStatusFailed  CharacterStatus = "failed"
	CharacterStatusSuccess CharacterStatus = "success"
)

type Result struct {
	ID           string          `json:"userCharacterId"`
	CharacterId  string          `json:"characterId"`
	Name         string          `json:"name"`
	ImgUrl       string          `json:"imgUrl"`
	Status       CharacterStatus `json:"status"`
	EmissionRate float64         `json:"emissionRate"`
}

func NewUserCharacter(userId string, characterId string, imgUrl string, emissionRate float64) (*UserCharacter, error) {
	now := flextime.Now()
	return &UserCharacter{
		UserId:       userId,
		CharacterId:  characterId,
		ImgUrl:       imgUrl,
		EmissionRate: emissionRate,
		Status:       CharacterStatusPending,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

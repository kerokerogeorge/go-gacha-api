package model

import (
	"time"

	"github.com/Songmu/flextime"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Result struct {
	ID          string    `json:"id" gorm:"primary_key"`
	UserId      string    `json:"userId"`
	CharacterId string    `json:"characterId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewResult(userId string, characterId string) (*Result, error) {
	now := flextime.Now()
	return &Result{
		UserId:      userId,
		CharacterId: characterId,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

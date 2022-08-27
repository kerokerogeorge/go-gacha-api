package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Result struct {
	ID          string    `json:"id" gorm:"primary_key"`
	UserId      string    `json:"userId"`
	CharacterId string    `json:"characterId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

package model

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID    string `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Token string `json:"token"`
	// Characters []*Character `json:"characters"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

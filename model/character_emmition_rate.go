package model

import "time"

type CharacterEmmitionRate struct {
	ID           string    `json:"id"`
	GachaID      string    `json:"gachaId"`
	CharacterID  string    `json:"characterId"`
	EmissionRate int       `json:"emissionRate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

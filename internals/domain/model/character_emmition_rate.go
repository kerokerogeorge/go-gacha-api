package model

import (
	"math/rand"
	"time"

	"github.com/Songmu/flextime"
)

type CharacterEmmitionRate struct {
	ID           string    `json:"id"`
	GachaID      string    `json:"gachaId"`
	CharacterID  string    `json:"characterId"`
	EmissionRate int       `json:"emissionRate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type CharacterWithEmmitionRate struct {
	CharacterID  string `json:"characterId"`
	Name         string `json:"name"`
	EmissionRate int    `json:"emissionRate"`
}

func NewCharacterEmmitionRate(gachaId string, characterId string) (*CharacterEmmitionRate, error) {
	rand.Seed(time.Now().UnixNano())
	emissionRate := rand.Intn(100-1) + 1
	now := flextime.Now()
	return &CharacterEmmitionRate{
		GachaID:      gachaId,
		CharacterID:  characterId,
		EmissionRate: emissionRate,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}

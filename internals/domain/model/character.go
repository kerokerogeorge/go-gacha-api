package model

import (
	"time"

	"github.com/Songmu/flextime"
)

type Character struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewCharacter(name string) (*Character, error) {
	now := flextime.Now()
	return &Character{
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

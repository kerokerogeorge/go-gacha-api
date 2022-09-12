package model

import (
	"time"

	"github.com/Songmu/flextime"
	"github.com/kerokerogeorge/go-gacha-api/internals/helper"
)

type Gacha struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewGacha() (*Gacha, error) {
	now := flextime.Now()
	return &Gacha{
		ID:        helper.NewULID().String(),
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

package model

import "time"

type Character struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	EmissionRate int       `json:"emissionRate"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

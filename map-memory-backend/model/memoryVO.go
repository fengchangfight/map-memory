package model

import "time"

type MemoryVO struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Longitude  float64   `json:"longitude"`
	Latitude   float64   `json:"latitude"`
	Icon       string    `json:"icon"`
	CreatedAt  time.Time `json:"created_at"`
	LastUpdate time.Time `json:"last_update"`
	Locked     bool      `json:"locked"`
	IsPublic   bool      `json:"is_public"`
}

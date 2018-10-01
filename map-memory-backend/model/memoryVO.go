package model

import "time"

type MemoryVO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
	Locked    bool      `json:"locked"`
	IsPublic  bool      `json:"is_public"`
}

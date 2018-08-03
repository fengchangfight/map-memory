package model

import "time"

type MemoryVO struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"`
}

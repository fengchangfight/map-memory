package model

type FavoriteLocVO struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

package model

type MemoryDetailVO struct {
	ID        int64   `json:"id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Icon      string  `json:"icon"`
	Username  string  `json:"username"`
	Nickname  string  `json:"nickname"`
	CreatedAt string  `json:"created_at"`
}

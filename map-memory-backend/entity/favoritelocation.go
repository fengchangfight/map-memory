package entity

import "time"

type FavoriteLocation struct {
	ID        int64 `gorm:"primary_key"`
	Name      string
	Longitude float64
	Latitude  float64
	User      User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID    int64
	LastUsed  time.Time
}

func (FavoriteLocation) TableName() string {
	return "mp_favorite_location"
}

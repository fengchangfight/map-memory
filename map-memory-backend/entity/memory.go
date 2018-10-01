package entity

import "time"

type Memory struct {
	ID        int64 `gorm:"primary_key"`
	Title     string
	Content   string
	Longitude float64
	Latitude  float64
	Icon      string
	User      User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	UserID    int64
	CreatedAt time.Time
	Locked    bool
	Openness  int //是否公开 0否 1是
}

func (Memory) TableName() string {
	return "mp_memory"
}

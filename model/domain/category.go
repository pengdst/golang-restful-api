package domain

import "time"

type Category struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Category) TableName() string {
	return "category"
}

package model

import "time"

type Base struct {
	Id        int `json:"id" column:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

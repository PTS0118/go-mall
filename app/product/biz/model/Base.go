package model

import "time"

type Base struct {
	Id        int32 `json:"id" column:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

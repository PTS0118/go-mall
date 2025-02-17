package model

import "time"

type Base struct {
	ID          int `gorm:"primarykey"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Base struct {
	Id        int                   `json:"id" column:"id" gorm:"primarykey"`
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import (
	"time"
)

type BaseModel struct {
	Id uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`

	CreatedAt time.Time `gorm:"column:created_at;index;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;not null" json:"updated_at"`
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	LanguageId uint64 `gorm:"primaryKey"`
	Language   string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

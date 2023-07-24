package entity

import (
	"time"

	"gorm.io/gorm"
)

type (
	Word struct {
		Word            string
		WordClass       string
		WordDescription string
		ExamplePhrase   string
	}

	IndonesianWord struct {
		IndonesianWordId uint64 `gorm:"primaryKey"`
		Word
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}

	EnglishWord struct {
		EnglishWordId uint64 `gorm:"primaryKey"`
		Word
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)

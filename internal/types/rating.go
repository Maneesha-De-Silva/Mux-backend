package types

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition


type Rating struct {
	RID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	
	RatingLevel      int    `json:"price"`
	Comment  string `json:"description"`
	CID uint
	ID uint
}





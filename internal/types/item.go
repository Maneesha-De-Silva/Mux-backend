package types

import (
	"time"

	"gorm.io/gorm"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Item struct {
	Model
	Title        string `json:"title"` // product Name
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`//quantity
	ReorderLevel int    `json:"reorder_level"`
	Description  string `json:"description"`
	CoverSrc     string `json:"coverSrc"`
	Rating []Rating
}

// Item name, price, quantity, reorder level, description

type Purcahse struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
	
}



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
	Title       string  `json:"title"`
	ServiceTime string  `json:"serviceTime"`
	DeliveryFee float64 `json:"deliveryFee"`
	Category    string  `json:"category"`
	Cuisine     string  `json:"cuisine"`
	Rating      int     `json:"rating"`
	Price       int     `json:"price"`
	CoverSrc    string  `json:"coverSrc"`
	Stock       int     `json:"stock"`
}

type Purcahse struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

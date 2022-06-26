package types

import "gorm.io/gorm"

type Item struct {
	gorm.Model
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

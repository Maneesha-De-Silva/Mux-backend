package types

import (
	//"time"

	"gorm.io/gorm"
)


type Order struct {
	gorm.Model
	DeliveryAddress    string    `json:"deliveryaddress"`
	Payment []Payment
	CID uint
}

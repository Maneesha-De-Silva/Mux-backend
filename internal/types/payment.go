package types

import (
//	"time"

	"gorm.io/gorm"
)


type Payment struct {
	gorm.Model
	PID        uint `gorm:"primaryKey" json:"id"`
	Amount  int       `json:"amount"`
	ID uint

}

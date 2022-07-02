package types

import (
	"time"

	"gorm.io/gorm"
)


type Customer struct {
	CID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	FName    string    `json:"fname"`
	LName    string    `json:"lname"`
	Title    string    `json:"title"`
	Address  string    `json:"address"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"dob"`
	Order [] Order
	Rating []Rating
}

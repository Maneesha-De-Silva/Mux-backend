package types

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	EID      uint      `gorm:"primaryKey" json:"id"`
	FName    string    `json:"fname"`
	LName    string    `json:"lname"`
	Title    string    `json:"title"`
	Address  string    `json:"address"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"dob"`
	Salary   int       `json:"salary"`
	Position string    `json:"position"`
}

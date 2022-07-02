package database

import (
	"backend/internal/types"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&types.Item{},
		&types.Purcahse{},
		&types.Employee{},
		&types.Payment{},
		&types.Rating{},
	); err != nil {
		return err
	}
	return nil
}

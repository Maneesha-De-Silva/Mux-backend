package database

import (
	"backend/internal/types"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&types.Item{}, &types.Purcahse{}); err != nil {
		return err
	}
	return nil
}

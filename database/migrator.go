package database

import (
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Profile{},
		&entity.Education{},
	)
	return err
}

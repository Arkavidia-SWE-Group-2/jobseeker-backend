package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type BaseEntity struct {
	ID string `json:"id" gorm:"primary_key"`
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if e.ID == "" {
		e.ID = ulid.Make().String()
	}

	return nil
}

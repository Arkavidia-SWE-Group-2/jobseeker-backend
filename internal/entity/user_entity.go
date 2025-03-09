package entity

import (
	"jobseeker/internal/pkg/helper"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Email    string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Phone    string    `json:"phone" gorm:"type:varchar(255);not null;unique"`
	Password string    `json:"password" gorm:"type:varchar(255);not null"`
	Vanity   string    `json:"vanity" gorm:"type:varchar(255);unique;not null"`
	Timestamp

	Profile *Profile `json:"profile,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	hashed, err := helper.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashed

	return nil
}

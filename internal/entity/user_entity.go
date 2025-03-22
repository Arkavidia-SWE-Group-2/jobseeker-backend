package entity

import (
	"jobseeker/internal/pkg/helper"

	"gorm.io/gorm"
)

type User struct {
	BaseEntity
	Email    string `json:"email" gorm:"column:email;not null;unique"`
	Phone    string `json:"phone" gorm:"column:phone;not null;unique"`
	Password string `json:"password" gorm:"column:password;not null"`
	Vanity   string `json:"vanity" gorm:"column:vanity;not null;unique"`
	Timestamp

	Profile   *Profile   `json:"profile,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Education *Education `json:"education,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

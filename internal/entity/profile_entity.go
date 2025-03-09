package entity

import (
	"github.com/google/uuid"
)

type Profile struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;unique"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Photo     string    `json:"photo"`
	Headline  string    `json:"headline"`
	About     string    `json:"about"`
	Timestamp

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *Profile) TableName() string {
	return "profiles"
}

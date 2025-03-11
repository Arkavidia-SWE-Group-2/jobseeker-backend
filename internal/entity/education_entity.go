package entity

import "time"

type Education struct {
	BaseEntity
	UserID      string    `json:"user_id" gorm:"column:user_id;not null;unique"`
	School      string    `json:"school" gorm:"column:school;not null"`
	Degree      string    `json:"degree" gorm:"column:degree"`
	Description string    `json:"description" gorm:"column:description"`
	StartDate   time.Time `json:"start_date" gorm:"column:start_date;type:date;not null"`
	EndDate     time.Time `json:"end_date" gorm:"column:end_date;type:date;not null"`
	Timestamp

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (e *Education) TableName() string {
	return "educations"
}

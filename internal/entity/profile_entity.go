package entity

type Profile struct {
	BaseEntity
	UserID    string `json:"user_id" gorm:"column:user_id;not null;unique"`
	FirstName string `json:"first_name" gorm:"column:first_name;not null"`
	LastName  string `json:"last_name" gorm:"column:last_name;not null"`
	Photo     string `json:"photo" gorm:"column:photo"`
	Headline  string `json:"headline" gorm:"column:headline"`
	About     string `json:"about" gorm:"column:about"`
	Timestamp

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *Profile) TableName() string {
	return "profiles"
}

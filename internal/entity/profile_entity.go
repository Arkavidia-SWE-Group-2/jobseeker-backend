package entity

type Profile struct {
	BaseEntity
	UserID    string `json:"user_id" gorm:"type:uuid;not null;unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo     string `json:"photo"`
	Headline  string `json:"headline"`
	About     string `json:"about"`
	Timestamp

	User *User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *Profile) TableName() string {
	return "profiles"
}

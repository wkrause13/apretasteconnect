package model

import "time"

// User is the generic user object
type User struct {
	ID        int        `json:"id"`
	Email     string     `json:"email" sql:"type:text;unique" valid:"email,length(0|500)"`
	Password  string     `json:"password" sql:"type:text" valid:"length(5|200)"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
}

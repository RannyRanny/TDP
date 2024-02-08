package model

import "time"

type TelegramAuthData struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Username  string     `json:"username"`
	PhotoURL  string     `json:"photo_url"`
	AuthDate  int64      `json:"auth_date"`
	Hash      string     `json:"hash"`
}

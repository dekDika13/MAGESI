package model

import "time"

type Users struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Roles     Roles     `gorm:"foreignKey:RoleId"`
	RoleId    uint      `json:"role_id"`
	Username  string    `gorm:"size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAT time.Time `json:"created_at"`
}

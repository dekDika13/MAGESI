package model

type Roles struct {
	RoleID   uint   `gorm:"primaryKey;autoIncrement" json:"role_id"`
	RoleName string `gorm:"size:50;not null" json:"role_name"`
}

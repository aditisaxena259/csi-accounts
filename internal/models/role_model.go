package models

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name string    `gorm:"type:text;not null;default:'external';check:name IN ('external', 'board', 'senior', 'junior')" json:"name"`
}

type RolePermission struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	RoleID       uuid.UUID `gorm:"type:uuid;not null" json:"roleID"`
	Permissions []Permission `gorm:"foreignKey:RoleID" json:"permissions"`
}
package schemas

import (
	"time"
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,max=50"`
	DateCreated time.Time `json:"dateCreated" validate:"required"`
}

type Permission struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,max=50"`
	Description string    `json:"description" validate:"max=255"`
}

type RolePermission struct {
	ID           uuid.UUID   `json:"id" validate:"required,uuid"`
	RoleID       uuid.UUID   `json:"roleID" validate:"required,uuid"`
	Permissions  []Permission `json:"permissions" validate:"required,dive"`
	DateAssigned time.Time   `json:"dateAssigned" validate:"required"`
}

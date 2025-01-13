package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Scope struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string    `json:"name" validate:"required,max=50"`
	Description string    `json:"description" validate:"max=255"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

type ClientScope struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	ClientID  uuid.UUID `json:"clientID" validate:"required,uuid"`
	ScopeID   uuid.UUID `json:"scopeID" validate:"required,uuid"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

type UserScope struct {
	ID        uuid.UUID `json:"id" validate:"required,uuid"`
	UserID    uuid.UUID `json:"userID" validate:"required,uuid"`
	ScopeID   uuid.UUID `json:"scopeID" validate:"required,uuid"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
}

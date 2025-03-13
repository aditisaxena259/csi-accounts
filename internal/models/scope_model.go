package models

import (
	"time"

	"github.com/google/uuid"
)

type Scope struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name        string    `gorm:"type:text;not null;check:name IN ('READ','WRITE','DELETE','UPDATE')" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp" json:"updatedAt"`
}

type ClientScope struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
    ClientID uuid.UUID `gorm:"type:uuid;not null;index" json:"clientID"`
    Client   Client    `gorm:"foreignKey:ClientID" json:"client"`
    ScopeID  uuid.UUID `gorm:"type:uuid;not null;index" json:"scopeID"`
    Scope    Scope     `gorm:"foreignKey:ScopeID" json:"scope"`
    CreatedAt time.Time `gorm:"default:current_timestamp" json:"createdAt"`
}

type UserScope struct {
    ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
    UserID   uuid.UUID `gorm:"type:uuid;not null;index" json:"userID"`
    User     User      `gorm:"foreignKey:UserID" json:"user"`
    ScopeID  uuid.UUID `gorm:"type:uuid;not null;index" json:"scopeID"`
    Scope    Scope     `gorm:"foreignKey:ScopeID" json:"scope"`
    CreatedAt time.Time `gorm:"default:current_timestamp" json:"createdAt"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name             string            `gorm:"type:text;not null;default:'external';check: Name IN ('external', 'board', 'senior', 'junior')" json:"name"`
	Date             time.Time         `gorm:"" json:"date"`
	Location         string            `gorm:"type:text" json:"location"`
	CreatedAt        time.Time         `gorm:"default:current_timestamp" json:"createdAt"`
	EventMemberships []EventMembership `gorm:"foreignKey:EventID" json:"-"`
}

type EventMembership struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	EventID   uuid.UUID `gorm:"type:uuid;not null" json:"eventID"`
	Event     Event     `gorm:"" json:"event"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	User      User      `gorm:"" json:"user"`
	Role      EventRole `gorm:"type:text;not null" json:"role"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"createdAt"`
}

type EventRole string

const (
	Participant EventRole = "participant"
	Coordinator EventRole = "coordinator"
	Admin       EventRole = "admin"
)

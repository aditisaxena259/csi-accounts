package schemas

import (
	"github.com/google/uuid"
	"time"
)


type CreateAuditLogBody struct {
	UserID    uuid.UUID `json:"userID" validate:"required"`
	ClientID  uuid.UUID `json:"clientID" validate:"required"`
	Action    string    `json:"action" validate:"required"`
	Timestamp time.Time `json:"auditLog" validate:"required"`
}


type UpdateAuditLogBody struct {
	UserID    *uuid.UUID `json:"userID"`
	ClientID  *uuid.UUID `json:"clientID"`
	Action    *string    `json:"action"`
	Timestamp *time.Time `json:"auditLog"`
}

package schemas

import "github.com/google/uuid"

// UpdateUserBody represents the structure for updating user details.
type UpdateUserBody struct {
	Name     string    `json:"name,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	RoleID   uuid.UUID `json:"roleID,omitempty"`
	Status   string    `json:"status,omitempty"` // Add Status here
}

package schemas

import "github.com/google/uuid"

// CreateUserBody defines the structure for the user creation request body
type CreateUserBody struct {
	Name         string    `json:"name" validate:"required"`
	Username     string    `json:"username" validate:"required,alphanum"`
	Email        string    `json:"email" validate:"required,email"`
	PasswordHash string    `json:"passwordHash" validate:"required"`
	RoleID       uuid.UUID `json:"roleID" validate:"required"`
}

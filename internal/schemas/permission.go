package schemas

type CreatePermissionBody struct {
	Name string `json:"name" validate:"required"`
}

// UpdatePermissionBody defines the request structure for updating a permission
type UpdatePermissionBody struct {
	Name *string `json:"name"`
}

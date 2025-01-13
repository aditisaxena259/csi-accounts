package schemas

import (

	"github.com/lib/pq"
)

// CreateClientBody defines the request structure for creating a client
type CreateClientBody struct {
	Name         string         `json:"name" validate:"required"`
	Description  string         `json:"description"`
	RedirectURIs pq.StringArray `json:"redirectURIs" validate:"required"`
	ClientScopes []string       `json:"clientScopes"`
	ClientID     string         `json:"clientID" validate:"required"`
	ClientSecret string         `json:"clientSecret" validate:"required"`
}

// UpdateClientBody defines the request structure for updating a client
type UpdateClientBody struct {
	Name         *string         `json:"name"`
	Description  *string         `json:"description"`
	RedirectURIs *pq.StringArray `json:"redirectURIs"`
	ClientScopes *[]string       `json:"clientScopes"`
	ClientID     *string         `json:"clientID"`
	ClientSecret *string         `json:"clientSecret"`
}

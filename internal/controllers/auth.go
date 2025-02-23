package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

// Struct to send JSON response
type AuthResponse struct {
	AuthCode string `json:"auth_code"`
}

// GenerateAuthorizationCode creates a random auth code
func GenerateAuthorizationCode() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// AuthHandler sends the auth code as JSON response
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	authCode := GenerateAuthorizationCode()
	response := AuthResponse{AuthCode: authCode}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

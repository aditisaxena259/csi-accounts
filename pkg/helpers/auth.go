package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"log"
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
		log.Println("Error generating random bytes:", err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// AuthHandler processes POST request and sends auth code as response
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure only POST requests are allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Decode incoming request (if needed)
	var requestData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Generate authorization code
	authCode := GenerateAuthorizationCode()
	if authCode == "" {
		http.Error(w, "Failed to generate authorization code", http.StatusInternalServerError)
		return
	}

	// Send the authorization code in the POST response
	response := AuthResponse{AuthCode: authCode}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	log.Println("✅ Authorization code generated and sent:", authCode)
}

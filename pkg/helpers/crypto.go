package helpers

import (
	"crypto/rand"
	"encoding/hex"
)


func GenerateRandomString(length int) string {
	bytes := make([]byte, length/2) 
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

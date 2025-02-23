package initializers

import (
	"log"
	"fmt"

	"csi-accounts/internal/models"
)

func RunMigrations() {
	if DB == nil {
		fmt.Println("Database not connected")
	}

    err := DB.AutoMigrate(
		&models.Client{}, 
		&models.Role{},          
		&models.Permission{},
		&models.Scope{},         // ✅ Scopes table must be created before ClientScope
		&models.User{},
		&models.Event{},
		&models.EventMembership{},
		&models.ClientScope{},   // ✅ Now, ClientScope can reference clients & scopes
		&models.UserScope{},
	)

	if err != nil {
		log.Fatal("Failed to run migrations")
	}
}
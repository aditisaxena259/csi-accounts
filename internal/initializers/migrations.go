package initializers

import (
	"log"

	"csi-accounts/internal/models"
	"gorm.io/gorm"
)

// RunMigrations applies database migrations safely.
func RunMigrations() {
	// 🔹 Check if DB is initialized before running migrations
	if DB == nil {
		log.Fatal("❌ Database connection is nil! Check if ConnectToDB() was called before migrations.")
		return
	}

	log.Println("🚀 Running database migrations...")

	// 🔹 Start a transaction to ensure atomic migrations
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 🔹 Check if models exist before migrating
		modelsToMigrate := []interface{}{
			&models.Client{},
			&models.Role{},
			&models.Permission{},
			&models.Scope{},
			&models.User{},
			&models.Event{},
			&models.EventMembership{},
			&models.ClientScope{},
			&models.UserScope{},
		}

		for _, model := range modelsToMigrate {
			if model == nil {
				log.Fatal("❌ One of the models is nil! Check struct definitions.")
				return nil
			}
		}

		// 🔹 Perform migrations in the correct order
		for _, model := range modelsToMigrate {
			if err := tx.AutoMigrate(model); err != nil {
				return err
			}
		}

		return nil
	})

	// 🔹 Handle migration errors
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
		return
	}

	log.Println("✅ Migrations completed successfully!")
}

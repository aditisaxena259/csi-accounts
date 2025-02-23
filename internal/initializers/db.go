package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL") // ✅ Fetch from environment variable
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL environment variable is not set!")
		return
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database:", err)
	} else {
		log.Println("✅ Connected to database!")
	}
}

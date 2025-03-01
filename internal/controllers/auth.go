package controllers

import (
	"csi-accounts/internal/initializers"
	"csi-accounts/internal/models"
	"csi-accounts/pkg/helpers"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ✅ Define request struct
type SignupRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   string `json:"roleID"`
}

// ✅ Signup Controller (Registers a new user)
func Signup(c *fiber.Ctx) error {
	var req SignupRequest

	// ✅ Parse JSON request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// ✅ Ensure all required fields are present
	if req.Name == "" || req.Username == "" || req.Email == "" || req.Password == "" || req.RoleID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required"})
	}

	// ✅ Convert RoleID string to UUID
	roleUUID, err := uuid.Parse(req.RoleID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid RoleID format"})
	}

	// ✅ Check if the role exists
	var existingRole models.Role
	if err := initializers.DB.Where("id = ?", roleUUID).First(&existingRole).Error; err != nil {
		// ✅ If role does not exist, create a new one
		newRole := models.Role{
			ID:   roleUUID,
			Name: "DefaultRole", // You can change this to "Admin", "User", etc.
		}
		if err := initializers.DB.Create(&newRole).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create role"})
		}
	}

	// ✅ Check if the email already exists
	var existingUser models.User
	if err := initializers.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "User already exists"})
	}

	// ✅ Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	// ✅ Create the user record
	user := models.User{
		Name:         req.Name,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Status:       models.ACTIVE,
		RoleID:       roleUUID,
		CreatedAt:    time.Now(),
		External:     false,
	}

	// ✅ Insert into the database
	if err := initializers.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	authCode := helpers.GenerateAuthorizationCode()

	// ✅ Return success response with auth code
	return c.JSON(fiber.Map{
		"message":            "✅ Login successful",
		"status":             "success",
		"authorization_code": authCode,
	})
}


// ✅ Login Controller (Validates user & returns auth code)
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ✅ Login Controller
func Login(c *fiber.Ctx) error {
	var req LoginRequest // ✅ Declare req variable properly

	// ✅ Debugging: Print raw request body
	log.Println("🔍 Raw Request Body:", string(c.Body()))

	// ✅ Parse JSON request body
	if err := c.BodyParser(&req); err != nil {
		log.Println("❌ BodyParser Error:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// ✅ Check for missing fields
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email and password are required"})
	}

	// ✅ Find user by email
	var user models.User
	if err := initializers.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// ✅ Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// ✅ Generate authorization code upon successful login
	authCode := helpers.GenerateAuthorizationCode()

	// ✅ Return success response with auth code
	return c.JSON(fiber.Map{
		"message":            "✅ Login successful",
		"status":             "success",
		"authorization_code": authCode,
	})
}
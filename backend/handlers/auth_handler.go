package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/database"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/models"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// ... (Fungsi Register, Login, dan ForgotPassword tidak berubah) ...
func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Name     string `json:"name" validate:"required,min=3"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not hash password"})
	}
	user := models.User{
		Name: input.Name, Email: input.Email, Password: string(hashedPassword), Role: "player",
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not create user, email might already exist"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": "User successfully registered"})
}
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	userResponse := struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role,
	}
	return c.JSON(fiber.Map{
		"status": "success", "message": "Success login", "data": fiber.Map{"token": t, "user": userResponse},
	})
}
func ForgotPassword(c *fiber.Ctx) error {
	type ForgotPasswordInput struct {
		Email string `json:"email" validate:"required,email"`
	}
	var input ForgotPasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.JSON(fiber.Map{"status": "success", "message": "If your email is registered, you will receive a reset link."})
	}
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not generate token"})
	}
	token := hex.EncodeToString(tokenBytes)
	passwordReset := models.PasswordReset{
		Email: input.Email, Token: token, ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	database.DB.Create(&passwordReset)
	log.Printf("RESET TOKEN for %s: %s\n", input.Email, token)
	return c.JSON(fiber.Map{"status": "success", "message": "If your email is registered, you will receive a reset link."})
}

// FUNGSI BARU DI SINI
func ResetPassword(c *fiber.Ctx) error {
	type ResetPasswordInput struct {
		Token    string `json:"token" validate:"required"`
		Password string `json:"password" validate:"required,min=6"`
	}
	var input ResetPasswordInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input"})
	}
	if err := validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	// 1. Cari token di database
	var passwordReset models.PasswordReset
	if err := database.DB.Where("token = ?", input.Token).First(&passwordReset).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid token"})
	}

	// 2. Cek apakah token sudah kedaluwarsa
	if time.Now().After(passwordReset.ExpiresAt) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Token has expired"})
	}

	// 3. Cari user berdasarkan email dari token
	var user models.User
	if err := database.DB.Where("email = ?", passwordReset.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User not found"})
	}

	// 4. Hash password baru
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Could not hash new password"})
	}

	// 5. Update password user di database
	database.DB.Model(&user).Update("password", string(newHashedPassword))

	// 6. Hapus token yang sudah dipakai dari database
	database.DB.Delete(&passwordReset)

	return c.JSON(fiber.Map{"status": "success", "message": "Password has been reset successfully"})
}

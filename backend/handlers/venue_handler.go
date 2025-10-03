package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/database"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/models"
	"gorm.io/gorm"
)

// ... (Fungsi GetAllVenues, GetVenue, CreateVenue, UpdateVenue, DeleteVenue tidak berubah) ...
func GetAllVenues(c *fiber.Ctx) error {
	var venues []models.Venue
	result := database.DB.Find(&venues)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get venues", "data": result.Error})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Venues Found", "data": venues})
}
func GetVenue(c *fiber.Ctx) error {
	id := c.Params("id")
	var venue models.Venue
	result := database.DB.First(&venue, id)
	if result.Error == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Venue not found", "data": nil})
	}
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get venue", "data": result.Error})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Venue Found", "data": venue})
}
func CreateVenue(c *fiber.Ctx) error {
	venue := new(models.Venue)
	if err := c.BodyParser(venue); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	result := database.DB.Create(&venue)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create venue", "data": result.Error})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Venue Created", "data": venue})
}
func UpdateVenue(c *fiber.Ctx) error {
	id := c.Params("id")
	var venue models.Venue
	if err := database.DB.First(&venue, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Venue not found", "data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get venue", "data": err})
	}
	var updateData models.Venue
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	venue.Name = updateData.Name
	venue.Address = updateData.Address
	venue.Description = updateData.Description
	database.DB.Save(&venue)
	return c.JSON(fiber.Map{"status": "success", "message": "Venue Updated", "data": venue})
}
func DeleteVenue(c *fiber.Ctx) error {
	id := c.Params("id")
	var venue models.Venue
	if err := database.DB.First(&venue, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Venue not found", "data": nil})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get venue", "data": err})
	}
	if err := database.DB.Delete(&venue).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to delete venue", "data": err})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Venue successfully deleted"})
}

// FUNGSI BARU DI SINI
func UploadVenueImage(c *fiber.Ctx) error {
	id := c.Params("id")
	var venue models.Venue

	// 1. Cek apakah venue ada
	if err := database.DB.First(&venue, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Venue not found"})
	}

	// 2. Ambil file dari form request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Image upload failed"})
	}

	// 3. Buat nama file unik untuk mencegah nama yang sama
	//    Contoh: "gambar-asli.jpg" -> "uuid-unik-gambar-asli.jpg"
	uniqueId := uuid.New()
	filename := uniqueId.String() + "-" + file.Filename

	// 4. Tentukan path untuk menyimpan file
	//    Kita akan membuat folder "uploads" di direktori backend
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", filename))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to save file"})
	}

	// 5. Buat URL yang bisa diakses dari frontend
	//    Contoh: "http://localhost:8000/uploads/uuid-unik-gambar-asli.jpg"
	imageURL := fmt.Sprintf("%s/uploads/%s", c.BaseURL(), filename)

	// 6. Update kolom image_url di database
	database.DB.Model(&venue).Update("image_url", imageURL)

	return c.JSON(fiber.Map{"status": "success", "message": "Image uploaded successfully", "data": imageURL})
}

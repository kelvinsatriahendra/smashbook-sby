package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/database"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/models"
)

// CreateCourtForVenue menambahkan data lapangan baru untuk sebuah venue
func CreateCourtForVenue(c *fiber.Ctx) error {
	venueId := c.Params("id") // Ambil ID venue dari URL

	court := new(models.Court)
	if err := c.BodyParser(court); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Set VenueID untuk lapangan baru ini
	// (Kita perlu konversi string ke uint)
	var venue models.Venue
	if err := database.DB.First(&venue, venueId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Venue not found", "data": err})
	}
	court.VenueID = venue.ID

	// Simpan court ke database
	if err := database.DB.Create(&court).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create court", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Court Created", "data": court})
}

// GetCourtsForVenue mengambil semua data lapangan untuk sebuah venue
func GetCourtsForVenue(c *fiber.Ctx) error {
	venueId := c.Params("id")
	var courts []models.Court

	// Cari semua court dengan venue_id yang cocok
	if err := database.DB.Where("venue_id = ?", venueId).Find(&courts).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not get courts", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Courts Found", "data": courts})
}
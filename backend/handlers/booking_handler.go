package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/database"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/models"
	"gorm.io/gorm"
)

// ... (fungsi CreateBooking dan GetUserBookings tidak berubah)
func CreateBooking(c *fiber.Ctx) error {
	courtId := c.Params("id")
	userId := c.Locals("user_id").(float64)

	type BookingRequest struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}

	var request BookingRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	startTime, err := time.Parse(time.RFC3339, request.StartTime)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid start_time format", "data": err})
	}
	endTime, err := time.Parse(time.RFC3339, request.EndTime)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid end_time format", "data": err})
	}

	var existingBooking models.Booking
	err = database.DB.Where(
		"court_id = ? AND status = 'confirmed' AND start_time < ? AND end_time > ?",
		courtId,
		endTime,
		startTime,
	).First(&existingBooking).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(409).JSON(fiber.Map{"status": "error", "message": "Schedule is already booked at this time"})
	}

	var court models.Court
	if err := database.DB.First(&court, courtId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Court not found"})
	}

	booking := models.Booking{
		CourtID:   court.ID,
		UserID:    uint(userId),
		StartTime: startTime,
		EndTime:   endTime,
		Status:    "confirmed",
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create booking", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Booking Created", "data": booking})
}
func GetUserBookings(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(float64)
	var bookings []models.Booking

	if err := database.DB.Where("user_id = ?", uint(userId)).Find(&bookings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not retrieve bookings"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Bookings found", "data": bookings})
}


// UBAH FUNGSI INI
func GetBookingsForOwnerVenue(c *fiber.Ctx) error {
	ownerId := c.Locals("user_id").(float64)

	var venue models.Venue
	if err := database.DB.Where("owner_id = ?", uint(ownerId)).First(&venue).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "You do not own any venue"})
	}

	var courts []models.Court
	database.DB.Where("venue_id = ?", venue.ID).Find(&courts)

	var courtIDs []uint
	for _, court := range courts {
		courtIDs = append(courtIDs, court.ID)
	}

	var bookings []models.Booking
	// GUNAKAN PRELOAD UNTUK MENGAMBIL DATA USER DAN COURT
	database.DB.Preload("User").Preload("Court").Where("court_id IN ?", courtIDs).Order("start_time desc").Find(&bookings)

	return c.JSON(fiber.Map{"status": "success", "message": "Bookings for your venue found", "data": bookings})
}
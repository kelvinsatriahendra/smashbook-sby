package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/database"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/handlers"
	"github.com/kelvinsatriahendra/smashbook-sby/backend/middleware"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	app.Use(cors.New())

	// MEMBUAT FOLDER UPLOADS MENJADI PUBLIK
	app.Static("/uploads", "./uploads")

	api := app.Group("/api")

	// ... (route lain tidak berubah)
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	api.Post("/forgot-password", handlers.ForgotPassword)
	api.Post("/reset-password", handlers.ResetPassword)
	api.Get("/venues", handlers.GetAllVenues)
	api.Get("/venues/:id", handlers.GetVenue)
	api.Get("/venues/:id/courts", handlers.GetCourtsForVenue)
	api.Post("/courts/:id/bookings", middleware.AuthMiddleware, handlers.CreateBooking)
	api.Get("/my-bookings", middleware.AuthMiddleware, handlers.GetUserBookings)
	api.Post("/venues", middleware.AuthMiddleware, handlers.CreateVenue)
	api.Put("/venues/:id", middleware.AuthMiddleware, handlers.UpdateVenue)
	api.Delete("/venues/:id", middleware.AuthMiddleware, handlers.DeleteVenue)
	api.Post("/venues/:id/courts", middleware.AuthMiddleware, handlers.CreateCourtForVenue)
	api.Get("/my-venue/bookings", middleware.AuthMiddleware, handlers.GetBookingsForOwnerVenue)

	// ROUTE BARU UNTUK UPLOAD GAMBAR
	api.Post("/venues/:id/upload-image", middleware.AuthMiddleware, handlers.UploadVenueImage)

	log.Fatal(app.Listen(":8000"))
}

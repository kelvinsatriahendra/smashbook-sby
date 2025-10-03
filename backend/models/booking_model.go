package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking merepresentasikan tabel 'bookings' di database
type Booking struct {
	gorm.Model
	CourtID   uint      `json:"court_id"`
	Court     Court     `gorm:"foreignKey:CourtID;references:ID" json:"court"` // Perbaikan di sini
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"`   // Perbaikan di sini
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `gorm:"varchar(50)" json:"status"`
}
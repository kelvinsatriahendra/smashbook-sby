package models

import (
	"time"

	"gorm.io/gorm" // <-- Pastikan gorm di-impor
)

type PasswordReset struct {
	gorm.Model      // <-- TAMBAHKAN BARIS INI
	Email     string    `gorm:"index"`
	Token     string    `gorm:"uniqueIndex"`
	ExpiresAt time.Time
}
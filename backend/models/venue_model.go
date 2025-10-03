package models

import "gorm.io/gorm"

// Venue merepresentasikan tabel 'venues' di database
type Venue struct {
	gorm.Model
	OwnerID     uint   `json:"owner_id"`
	Name        string `gorm:"varchar(255)" json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"`
	ImageURL    string `gorm:"default:null" json:"image_url"` // <-- TAMBAHKAN BARIS INI
}
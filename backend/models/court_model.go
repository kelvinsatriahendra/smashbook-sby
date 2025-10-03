package models

import "gorm.io/gorm"

type Court struct {
	gorm.Model
	VenueID   uint   `json:"venue_id"`
	Name      string `gorm:"varchar(100)" json:"name"`
	FloorType string `gorm:"varchar(50)" json:"floor_type"`
}

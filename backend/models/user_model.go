package models

import "gorm.io/gorm"

// User merepresentasikan tabel 'users' di database
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	// TAMBAHKAN FIELD INI JIKA BELUM ADA
	Role string `gorm:"varchar(50)" json:"role"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey" json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	User_Id   int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      User       `gorm:"foreignKey:User_Id" json:"user"`
}

type PhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoList struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoURL  string     `json:"photo_url"`
	User_Id   int        `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      PhotoUser  `json:"user"`
}

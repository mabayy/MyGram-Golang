package models

import (
	"time"

	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"type:varchar(100)" json:"name"`
	SocialMediaURL string     `gorm:"type:text" json:"social_media_url"`
	User_Id        int        `json:"user_id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	User           User       `gorm:"foreignKey:User_Id" json:"user"`
}

type SocialMediaUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SocialMediaList struct {
	ID             uint            `json:"id"`
	Name           string          `gorm:"type:varchar(100)" json:"name"`
	SocialMediaURL string          `gorm:"type:text" json:"social_media_url"`
	User_Id        int             `json:"user_id"`
	CreatedAt      *time.Time      `json:"created_at"`
	UpdatedAt      *time.Time      `json:"updated_at"`
	User           SocialMediaUser `json:"user"`
}

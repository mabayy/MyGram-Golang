package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey" json:"id"`
	User_Id   int        `json:"user_id"`
	Photo_Id  int        `json:"photo_id"`
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Photo     Photo      `gorm:"foreignKey:Photo_Id" json:"photo"`
	User      User       `gorm:"foreignKey:User_Id" json:"user"`
}

type CommentUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	User_Id  int    `json:"user_id"`
}

type CommentList struct {
	ID        uint         `json:"id"`
	User_Id   int          `json:"user_id"`
	Photo_Id  int          `json:"photo_id"`
	Message   string       `json:"message"`
	CreatedAt *time.Time   `json:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at"`
	Photo     CommentPhoto `json:"photo"`
	User      CommentUser  `json:"user"`
}

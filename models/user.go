package models

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Age       int        `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	Password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		fmt.Println("Failed to encrypt password: ", err)
		return err
	}
	u.Password = string(Password)
	return nil
}

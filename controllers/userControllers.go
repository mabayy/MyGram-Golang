package controllers

import (
	"final-projek/helpers"
	"final-projek/models"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func (db *UserDB) UserRegister(c *gin.Context) {
	var (
		User    models.User
		NewUser models.User
	)
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if User.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}

	if User.Username == "" {
		c.JSON(400, gin.H{
			"message": "Username is required",
		})
		return
	}

	if User.Password == "" {
		c.JSON(400, gin.H{
			"message": "Password is required",
		})
		return
	}

	if len(User.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "Password must contain 6 characters",
		})
		return
	}

	if User.Age < 8 {
		c.JSON(400, gin.H{
			"message": "Minimun age is 8 years",
		})
		return
	}

	_, errMailFormat := mail.ParseAddress(User.Email)
	if errMailFormat != nil {
		c.JSON(400, gin.H{
			"message": "Email format is wrong",
		})
		return
	}

	db.DB.Where("email = ?", User.Email).First(&NewUser)
	if NewUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	db.DB.Where("username = ?", User.Username).First(&NewUser)
	if NewUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Username already used",
		})
		return
	}

	errCreate := db.DB.Debug().Create(&User).Error
	if errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(201, gin.H{
		"age":      User.Age,
		"email":    User.Email,
		"id":       User.ID,
		"username": User.Username,
	})
}

func (db *UserDB) UserLogin(c *gin.Context) {
	var User models.User

	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	dbResult := models.User{}
	errUser := db.DB.Debug().Where("email = ?", User.Email).Last(&dbResult).Error
	if errUser != nil {
		c.AbortWithError(http.StatusInternalServerError, errUser)
		return
	}

	errBcrypt := bcrypt.CompareHashAndPassword([]byte(dbResult.Password), []byte(User.Password))
	if errBcrypt != nil {
		c.AbortWithError(http.StatusBadRequest, errBcrypt)
		return
	}

	token := helpers.GenerateToken(dbResult.Username)

	c.JSON(200, gin.H{
		"token": token,
	})
}

func (db *UserDB) UserUpdate(c *gin.Context) {
	id := c.Param("userId")
	userId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		fmt.Println("error found: ", errConvert)
		c.JSON(400, gin.H{
			"result": "params userId is required",
		})
		return
	}
	var (
		User    models.User
		NewUser models.User
	)

	errUser := db.DB.First(&User, userId).Error
	if errUser != nil {
		c.JSON(404, gin.H{
			"result": "Data not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&User); err != nil {
		fmt.Println("error found: ", err)
		c.JSON(400, gin.H{
			"result": "Bad Request",
		})
		return
	}

	if User.Email == "" {
		c.JSON(400, gin.H{
			"message": "Email is required",
		})
		return
	}

	if User.Username == "" {
		c.JSON(400, gin.H{
			"message": "Username is required",
		})
		return
	}

	_, errMailFormat := mail.ParseAddress(User.Email)
	if errMailFormat != nil {
		c.JSON(400, gin.H{
			"message": "Email format is warong",
		})
		return
	}

	db.DB.Where("email = ?", User.Email).First(&NewUser)
	if NewUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Email already used",
		})
		return
	}

	db.DB.Where("username = ?", User.Username).First(&NewUser)
	if NewUser != (models.User{}) {
		c.JSON(400, gin.H{
			"message": "Username already used",
		})
		return
	}

	errUpdate := db.DB.Model(&User).Updates(models.User{Username: User.Username, Email: User.Email}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"result": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"id":         User.ID,
		"email":      User.Email,
		"username":   User.Username,
		"age":        User.Age,
		"updated_at": User.UpdatedAt,
	})
}

func (db *UserDB) UserDelete(c *gin.Context) {
	var (
		User   models.User
		result gin.H
	)
	id := c.GetString("userId")
	err := db.DB.First(&User, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	} else {
		errDelete := db.DB.Delete(&User).Error
		if errDelete != nil {
			result = gin.H{
				"result": "delete failed",
			}
		} else {
			result = gin.H{
				"message": "Your account has been successfully deleted",
			}
		}
	}

	c.JSON(200, result)
}

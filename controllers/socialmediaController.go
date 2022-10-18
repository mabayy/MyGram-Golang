package controllers

import (
	"final-projek/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaDB struct {
	DB *gorm.DB
}

func (db *SocialMediaDB) CreateSocialMedia(c *gin.Context) {
	var sosmed models.SocialMedia
	userId := c.GetString("userId")
	userIdConvert, _ := strconv.Atoi(userId)
	sosmed.User_Id = userIdConvert

	err := c.ShouldBindJSON(&sosmed)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if sosmed.Name == "" {
		c.JSON(400, gin.H{
			"message": "Name is required",
		})
		return
	}

	if sosmed.SocialMediaURL == "" {
		c.JSON(400, gin.H{
			"message": "Social_Media_Url is required",
		})
		return
	}

	errCreate := db.DB.Debug().Create(&sosmed).Error
	if errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(201, gin.H{
		"id":               sosmed.ID,
		"name":             sosmed.Name,
		"social_media_url": sosmed.SocialMediaURL,
		"user_id":          sosmed.User_Id,
		"created_at":       sosmed.CreatedAt,
	})
}

func (db *SocialMediaDB) GetSocialMedias(c *gin.Context) {
	var (
		socialMedia    []models.SocialMedia
		socialMediaNew []models.SocialMediaList
	)

	db.DB.Preload("User").Find(&socialMedia)

	for _, list := range socialMedia {
		socialMediaNew = append(socialMediaNew, models.SocialMediaList{
			ID:             list.ID,
			Name:           list.Name,
			SocialMediaURL: list.SocialMediaURL,
			User_Id:        list.User_Id,
			CreatedAt:      list.CreatedAt,
			UpdatedAt:      list.UpdatedAt,
			User: models.SocialMediaUser{
				ID:       list.User.ID,
				Email:    list.User.Email,
				Username: list.User.Username,
			},
		})
	}

	c.JSON(200, gin.H{
		"social_medias": socialMediaNew,
	})
}

func (db *SocialMediaDB) UpdateSocialMedia(c *gin.Context) {
	id := c.Param("socialMediaId")
	socialMediaId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "params socialMediaId is required",
		})
		return
	}

	var socialMedia models.SocialMedia
	errComment := db.DB.First(&socialMedia, socialMediaId).Error
	if errComment != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	}

	req := models.SocialMedia{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Name == "" {
		c.JSON(400, gin.H{
			"message": "Name is required",
		})
		return
	}

	if req.SocialMediaURL == "" {
		c.JSON(400, gin.H{
			"message": "Social_Media_Url is required",
		})
		return
	}

	errUpdate := db.DB.Model(&socialMedia).Updates(models.SocialMedia{Name: req.Name, SocialMediaURL: req.SocialMediaURL}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(201, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaURL,
		"user_id":          socialMedia.User_Id,
		"created_at":       socialMedia.CreatedAt,
	})
}

func (db *SocialMediaDB) DeleteSocialMedia(c *gin.Context) {
	var (
		socialMedia models.SocialMedia
	)

	id := c.Param("socialMediaId")
	socialMediaId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := db.DB.First(&socialMedia, socialMediaId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	} else {
		errDelete := db.DB.Delete(&socialMedia).Error
		if errDelete != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "Your social media has been successfully deleted",
			})
			return
		}
	}
}

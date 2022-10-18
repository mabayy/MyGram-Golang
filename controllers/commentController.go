package controllers

import (
	"final-projek/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentDB struct {
	DB *gorm.DB
}

func (db *CommentDB) CreateComment(c *gin.Context) {
	var req models.Comment
	userId := c.GetString("userId")
	userIdConvert, _ := strconv.Atoi(userId)
	req.User_Id = userIdConvert

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Message == "" {
		c.JSON(400, gin.H{
			"message": "Message is required",
		})
		return
	}

	errCreate := db.DB.Debug().Create(&req).Error
	if errCreate != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(201, gin.H{
		"id":         req.ID,
		"message":    req.Message,
		"photo_id":   req.Photo_Id,
		"user_id":    req.User_Id,
		"created_at": req.CreatedAt,
	})
}

func (db *CommentDB) GetComments(c *gin.Context) {
	var (
		comments    []models.Comment
		commentsRes []models.CommentList
	)

	db.DB.Preload("User").Preload("Photo").Find(&comments)
	for _, list := range comments {
		commentsRes = append(commentsRes, models.CommentList{
			ID:        list.ID,
			Message:   list.Message,
			Photo_Id:  list.Photo_Id,
			User_Id:   list.User_Id,
			CreatedAt: list.CreatedAt,
			UpdatedAt: list.UpdatedAt,
			User: models.CommentUser{
				ID:       list.User.ID,
				Email:    list.User.Email,
				Username: list.User.Username,
			},
			Photo: models.CommentPhoto{
				ID:       list.Photo.ID,
				Title:    list.Photo.Title,
				Caption:  list.Photo.Caption,
				PhotoURL: list.Photo.PhotoURL,
				User_Id:  list.Photo.User_Id,
			},
		})
	}

	c.JSON(200, commentsRes)
}

func (db *CommentDB) UpdateComment(c *gin.Context) {
	id := c.Param("commentId")
	commentId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "params photoId is required",
		})
		return
	}

	var comment models.Comment
	errComment := db.DB.First(&comment, commentId).Error
	if errComment != nil {
		c.JSON(400, gin.H{
			"message": "Data not found",
		})
		return
	}

	req := models.Comment{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if req.Message == "" {
		c.JSON(400, gin.H{
			"message": "Message is required",
		})
		return
	}

	errUpdate := db.DB.Model(&comment).Updates(models.Comment{Message: req.Message}).Error
	if errUpdate != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(201, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.Photo_Id,
		"user_id":    comment.User_Id,
		"created_at": comment.CreatedAt,
	})
}

func (db *CommentDB) DeleteComment(c *gin.Context) {
	var (
		comment models.Comment
	)

	id := c.Param("commentId")
	commentId, errConvert := strconv.Atoi(id)
	if errConvert != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	err := db.DB.First(&comment, commentId).Error
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Data not found",
		})
		return
	} else {
		errDelete := db.DB.Delete(&comment).Error
		if errDelete != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		} else {
			c.JSON(200, gin.H{
				"message": "Your comment has been successfully deleted",
			})
			return
		}
	}
}

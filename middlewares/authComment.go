package middlewares

import (
	"errors"
	"final-projek/database"
	"final-projek/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("commentId")
		idConvert, errConvert := strconv.Atoi(id)
		if errConvert != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("bad request"))
			c.JSON(400, gin.H{
				"message": "params commentId is required",
			})
			return
		}

		result := models.Comment{}
		errFind := database.StartDB().First(&result, idConvert).Error
		if errFind != nil {
			c.AbortWithError(404, errors.New("data not found"))
			c.JSON(404, gin.H{
				"message": "Data not found",
			})
			return
		} else {
			userId := c.GetString("userId")
			userIdConvert, _ := strconv.Atoi(userId)
			if result.User_Id != userIdConvert {
				c.AbortWithError(403, errors.New("forbidden access"))
				c.JSON(404, gin.H{
					"message": "Forbidden access",
				})
				return
			} else {
				c.Next()
			}
		}
	}
}

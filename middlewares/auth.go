package middlewares

import (
	"errors"

	"final-projek/database"
	"final-projek/helpers"
	"final-projek/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization"))
			c.JSON(401, gin.H{
				"message": "Unautorized",
			})
			return
		}

		authSplit := strings.Split(auth, " ")
		if len(authSplit) != 2 {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization"))
			c.JSON(401, gin.H{
				"message": "Unautorized",
			})
			return
		}
		if authSplit[0] != "Bearer" {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization"))
			c.JSON(401, gin.H{
				"message": "Unautorized",
			})
			return
		}

		token, err := helpers.VerifyToken(authSplit[1])
		if err != nil {
			fmt.Println("Bad Authorization")
			c.AbortWithError(http.StatusBadRequest, errors.New("bad authorization"))
			c.JSON(401, gin.H{
				"message": "Unautorized",
			})
			return
		}

		dbResult := models.User{}
		username := token["username"]
		err = database.StartDB().Debug().Where("username = ?", username).Last(&dbResult).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Username not found")
				c.AbortWithError(http.StatusBadRequest, errors.New("username not found"))
				c.JSON(401, gin.H{
					"message": "Unautorized",
				})
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			c.JSON(401, gin.H{
				"message": "Unautorized",
			})
			return
		}
		c.Set("userId", strconv.FormatInt(int64(dbResult.ID), 10))
		c.Next()
	}
}

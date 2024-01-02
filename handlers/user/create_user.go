package user

import (
	"fmt"
	database "go-api/db"
	CheckPassword "go-api/internal/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user database.UserRequest

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(user)
		user.Name = strings.TrimSpace(user.Name)

		if user.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tên người dùng không được để trống"})
			return
		}
		user.Password = strings.TrimSpace(user.Password)

		if user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu không được để trống"})
			return
		}
		newUser := database.User{
			Name:     user.Name,
			Password: CheckPassword.HashPassword(user.Password),
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

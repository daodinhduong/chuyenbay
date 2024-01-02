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

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request database.UserRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(request)
		request.Name = strings.TrimSpace(request.Name)

		if request.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tên người dùng không được để trống"})
			return
		}
		request.Password = strings.TrimSpace(request.Password)

		if request.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu không được để trống"})
			return
		}
		user, err := database.GetUserByEmail(c.Request.Context(), db, request.Name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		credentialError := CheckPassword.CheckPassword(request.Password, user)
		if credentialError != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Sai mật khẩu",
			})
			return
		}
		tokenString, err := CheckPassword.GenerateJWT(request.Name, request.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

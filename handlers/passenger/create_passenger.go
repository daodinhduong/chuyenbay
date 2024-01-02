package passenger

import (
	"fmt"
	database "go-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Createpassenser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var passenser database.Passenger

		if err := c.ShouldBindJSON(&passenser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(passenser)
		passenser.Name = strings.TrimSpace(passenser.Name)

		if passenser.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tên khách hàng không được để trống"})
			return
		}
		passenser.Gender = strings.TrimSpace(passenser.Gender)

		if passenser.Gender == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Giới tính không được để trống"})
			return
		}

		if passenser.Age == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tuổi không được để trống"})
			return
		}

		if err := db.Create(&passenser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": passenser})
	}
}

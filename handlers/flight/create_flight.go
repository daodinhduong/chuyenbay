package flight

import (
	database "go-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFlight(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var flight database.Flights

		if err := c.ShouldBindJSON(&flight); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		flight.Name = strings.TrimSpace(flight.Name)

		if flight.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tên chuyến bay không được để trống"})
			return
		}
		flight.From = strings.TrimSpace(flight.From)

		if flight.From == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Điểm bắt đầu không được để trống"})
			return
		}
		// flight.To = strings.TrimSpace(flight.To)

		// if flight.To == "" {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Điểm đến không được để trống"})
		// 	return
		// }

		flight.Status = 1
		if err := db.Create(&flight).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": flight})
	}
}

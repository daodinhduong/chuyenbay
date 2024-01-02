package flight

import (
	"errors"
	"github.com/gin-gonic/gin"
	database "go-api/db"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteFlightById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var existingFlight database.Flights
		if err := db.Table("flights").Where("id = ?", id).First(&existingFlight).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Bản ghi không tồn tại"})
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
		if err := db.Table("flights").
			Where("id = ?", id).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

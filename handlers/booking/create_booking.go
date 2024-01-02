package booking

import (
	database "go-api/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBooking(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var booking database.Booking

		if err := c.ShouldBindJSON(&booking); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&database.Flights{}).Where("id = ?", booking.FlightID).First(&database.Flights{}).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Chuyến bay không tồn tại"})
			return
		}
		if err := db.Model(&database.Passenger{}).Where("id = ?", booking.PassengerID).First(&database.Passenger{}).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hành khách không tồn tại"})
			return
		}
		booking.SeatNumber = strings.TrimSpace(booking.SeatNumber)

		if booking.SeatNumber == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Vị trí ngồi không được để trống"})
			return
		}

		if booking.FlightID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mã chuyến bay không được để trống"})
			return
		}
		if booking.PassengerID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mã khách hàng không được để trống"})
			return
		}

		if err := db.Create(&booking).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": booking})
	}
}

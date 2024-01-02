package passenger

import (
	"fmt"
	database "go-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfPassenger(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		type DataPaging struct {
			Page  int   `json:"page" form:"page"`
			Limit int   `json:"limit" form:"limit"`
			Total int64 `json:"total" form:"-"`
		}
		var paging DataPaging
		if err := c.ShouldBindJSON(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(paging)
		if paging.Page <= 0 {
			paging.Page = 1
		}

		if paging.Limit <= 0 {
			paging.Limit = 10
		}

		offset := (paging.Page - 1) * paging.Limit

		var result []database.Passenger

		if err := db.Model(&database.Passenger{}).Preload("Bookings").Preload("Bookings.Flight").
			Offset(offset).
			Limit(paging.Limit).
			Count(&paging.Total).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var res []database.PassengerRes
		for _, v := range result {
			var newBooking []database.BookingRest
			for _, t := range v.Bookings {
				flight := database.FlightRest{
					Name:   t.Flight.Name,
					From:   t.Flight.From,
					To:     t.Flight.To,
					Status: t.Flight.Status,
				}
				newBooking = append(newBooking, database.BookingRest{
					FlightID:   t.FlightID,
					TicketID:   t.TicketID,
					Flight:     flight,
					SeatNumber: t.SeatNumber,
				})
			}
			res = append(res, database.PassengerRes{
				Name:    v.Name,
				Gender:  v.Gender,
				Age:     v.Age,
				Booking: newBooking,
			})
		}
		newRes := database.SuccessResPassenger{
			Data:   res,
			Paging: paging,
		}
		c.JSON(http.StatusOK, gin.H{"data": newRes})
	}
}

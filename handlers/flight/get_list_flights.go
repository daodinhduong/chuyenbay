package flight

import (
	"fmt"
	database "go-api/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfFlight(db *gorm.DB) gin.HandlerFunc {
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

		var result []database.Flights

		if err := db.Table("flights").
			Offset(offset).
			Limit(paging.Limit).
			Count(&paging.Total).
			Order("id desc").
			Find(&result).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var newFlightResult []database.FlightRest
		for _, v := range result {
			newFlightResult = append(newFlightResult, database.FlightRest{
				Name:   v.Name,
				From:   v.From,
				To:     v.To,
				Status: v.Status,
			})
		}
		newRes := database.SuccessRes{
			Data:   newFlightResult,
			Paging: paging,
		}
		c.JSON(http.StatusOK, gin.H{"data": newRes})
	}
}

package main

import (
	database "go-api/db"
	booking "go-api/handlers/booking"
	flight "go-api/handlers/flight"
	passenger "go-api/handlers/passenger"
	user "go-api/handlers/user"
	"go-api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456789a@tcp(127.0.0.1:3306)/airport?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)
	db.AutoMigrate(&database.Flights{}, &database.Passenger{}, &database.Ticket{}, &database.Booking{}, &database.User{})
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		users := v1.Group("/user")
		{
			users.POST("/register", user.CreateUser(db))
			users.POST("/login", user.Login(db))
		}
		//chuyen bay
		v1.Use(middleware.Auth()).POST("/flights", flight.CreateFlight(db))
		v1.GET("/flights", flight.GetListOfFlight(db))
		v1.GET("/flights/:id", flight.GetFilghtByID(db))
		v1.PUT("/flights/:id", flight.EditFlightById(db))
		v1.DELETE("/flights/:id", flight.DeleteFlightById(db))
		//khach hang
		v1.POST("/passenger", passenger.Createpassenser(db))
		v1.GET("/passenger", passenger.GetListOfPassenger(db))
		v1.GET("/passenger/:id", passenger.GetPassengerByID(db))
		v1.PUT("/passenger/:id", passenger.DeletePassengerById(db))
		v1.DELETE("/passenger/:id", passenger.DeletePassengerById(db))
		//booking
		v1.POST("/booking", booking.CreateBooking(db))
		v1.GET("/booking", booking.GetListOfBooking(db))
		v1.GET("/booking/:id", booking.GetBookingByID(db))
		v1.PUT("/booking/:id", booking.DeleteBookingById(db))
		v1.DELETE("/booking/:id", booking.DeleteBookingById(db))
	}
	router.Run()
}

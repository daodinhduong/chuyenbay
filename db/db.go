package data

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Flights struct {
	gorm.Model
	Name      string     `json:"name" gorm:"column:name;"`
	From      string     `json:"from" gorm:"column:from;"`
	To        string     `json:"to" gorm:"column:to;"`
	StartTime *time.Time `json:"start_time" gorm:"column:start_time"`
	Status    uint       `json:"status" gorm:"column:status;"`
	Bookings  []Booking  `json:"bookings,omitempty" gorm:"foreignKey:FlightID"`
}

type Passenger struct {
	gorm.Model
	Name     string    `json:"name" gorm:"column:name;"`
	Gender   string    `json:"gender" gorm:"column:gender;"`
	Age      int       `json:"age" gorm:"column:age;"`
	Bookings []Booking `json:"bookings,omitempty" gorm:"foreignKey:PassengerID"`
}
type Booking struct {
	gorm.Model
	FlightID    uint    `json:"flight_id" gorm:"column:flight_id;"`
	Flight      Flights `json:"flight,omitempty" gorm:"foreignKey:FlightID"`
	PassengerID uint    `json:"passenger_id" gorm:"column:passenger_id;"`
	TicketID    uint    `json:"ticket_id" gorm:"column:ticket_id;"`
	SeatNumber  string  `json:"seat_number" gorm:"column:seat_number;"`
}
type Ticket struct {
	gorm.Model
	Name     string    `json:"name" gorm:"column:name;"`
	Bookings []Booking `json:"bookings,omitempty" gorm:"foreignKey:TicketID"`
}
type User struct {
	gorm.Model
	Name        string    `json:"name" gorm:"column:name;"`
	Password    []byte    `json:"password" gorm:"column:password;"`
	Passenger   Passenger `json:"passenger,omitempty" gorm:"foreignKey:PassengerID"`
	PassengerID *uint     `json:"passenger_id,omitempty" gorm:"column:passenger_id;"`
}
type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func GetUserByEmail(ctx context.Context, database *gorm.DB, name string) (*User, error) {
	var user User
	err := database.Model(&User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

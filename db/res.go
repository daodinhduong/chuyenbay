package data

type SuccessRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging"`
}
type SuccessResPassenger struct {
	Data   []PassengerRes `json:"data"`
	Paging interface{}    `json:"paging"`
}
type FlightRest struct {
	Name   string
	From   string
	To     string
	Status uint
}
type BookingRest struct {
	FlightID   uint
	TicketID   uint
	SeatNumber string
	Flight     FlightRest
}
type PassengerRes struct {
	Booking []BookingRest
	Name    string
	Age     int
	Gender  string
}

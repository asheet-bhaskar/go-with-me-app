package service

type Services struct {
	Booking BookingService
}

func NewServices() *Services {
	return &Services{
		Booking: *NewBookingService(),
	}
}

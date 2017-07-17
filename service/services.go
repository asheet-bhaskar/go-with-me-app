package service

type Services struct {
	Booking  BookingService
	Customer CustomerService
	Driver   DriverService
}

func NewServices() *Services {
	return &Services{
		Booking:  *NewBookingService(),
		Driver:   *NewDriverService(),
		Customer: *NewCustomerService(),
	}
}

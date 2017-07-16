package factories

import (
	"github.com/bluele/factory-go/factory"
	"github.com/heroku/go-with-me-app/domain"
)

var BookingFactory = factory.NewFactory(&domain.Booking{})

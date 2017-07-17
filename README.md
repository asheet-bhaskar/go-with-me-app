Problem:
when we book a ride, even after we get the driver we wait for driver to 
to come sometimes driver is very far it takes too long.
At the same time we see some driver around without booking, 
so instead of waiting for driver to come, why can't we 
take ride with driver near me.

Solution
We worked on QR code based solution. 
1- customer enters pickuop and destination
2- customer get the estimated fare
3- on customer a QR code is generated and customer approach to
   nearby driver.
4- driver scnns the QR code and booking get created on 
   successful scan of QR code.

Approach:
we started working from scratch and made simplest assumptions.
we worked on following three modules
1. customer App
2. driver app
3. backend service

Implementation details are as follows
1. customer App
	a) enter pickup and destintaion (we have used autocomplete API for this)
	b) get the estimate fare
	c) call create booking API, it will return booking_id
	d) using booking_id generate QR code.
	e) there is poll on get booking status API each 5 seconds for 2 minutes, if after two minutes QR sacn is not successfull then call set booking status to driver not found API. (not implemented)

2. Driver App
	a) scan the QR code and decode booking_id
	b) call start trip API. (not implemented)
	c) call complete trip API. (not implemented)

3. Backend service 
	we wrote a backend service in golang and use postgres as database.
	we have deployed this service on heroku.
	backend service has following APIs.
	a) create booking 
		path: "/v1/booking/create"
	b) get booking status
		path: "/v1/booking/status"
	c) set the booking status if driver not found after threshold time
		path: "/v1/booking/status/driver_not_found"
	d) get estimate_fare based on distance
		path: "/v1/booking/fare/estimate"
	e) create a customer
		path: "/v1/customer/create"
	f) create a driver
		path: "/v1/driver/create"
	g) update driver status
		path: "/v1/driver/status"
	h) get driver status
		path: "/v1/driver/status"
	i) start the trip
		path: "/v1/trip/start"
	j) complete the trip
		path: "/v1/trip/complete"
	
this service is deployed on heroku, host url is https://go-with-me-application.herokuapp.com


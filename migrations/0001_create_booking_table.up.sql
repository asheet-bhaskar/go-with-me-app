CREATE TABLE bookings(
	id bigserial primary key,
	booking_id varchar(20) not null,
	customer_id varchar ,
	driver_id varchar,
	pick_up varchar not null,
	destination varchar not null,
	fare varchar not null,
	status varchar not null,
	created_at timestamp without time zone not null, 
	updated_at timestamp without time zone not null 
);


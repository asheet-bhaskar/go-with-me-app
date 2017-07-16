CREATE TABLE customers(
	id bigserial primary key,
	driver_id varchar not null,
	name varchar not null,
	email varchar not null,
	status varchar not null,
	created_at timestamp without time zone not null, 
	updated_at timestamp without time zone not null 
);


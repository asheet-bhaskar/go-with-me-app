ALTER TABLE customers DROP COLUMN driver_id;
ALTER TABLE customers DROP COLUMN status;
ALTER TABLE  customers ADD COLUMN customer_id varchar;

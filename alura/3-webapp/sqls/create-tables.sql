CREATE TABLE products (
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	quantity integer
);

CREATE USER alura_store WITH PASSWORD 'ExamplePsw321.';

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO alura_store;

insert into products 
(name, description, price, quantity) values 
('Shirt', 'Blue', 29, 10),
('Notebook', 'Red', 1999, 2),
('Phone', 'Black', 199, 20);
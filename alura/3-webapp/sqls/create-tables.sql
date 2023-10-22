create table products (
	id serial primary key,
	name varchar,
	description varchar,
	price decimal,
	quantity integer
);

insert into products 
(name, description, price, quantity) values 
('Shirt', 'Blue', 29, 10),
('Notebook', 'Red', 1999, 2),
('Phone', 'Black', 199, 20);

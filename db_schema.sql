create table if not exists products (
	id SERIAL primary key,
	name VARCHAR(255) not null,
	price numeric(10, 2) not null,
	quantity numeric
);

delete from products;

insert into products(name, price, quantity) values
	('Macbook Pro', 1299.90, 5),
	('iPhone', 799.90, 10),
	('Airpods Max', 449.90, 3);

select * from products;
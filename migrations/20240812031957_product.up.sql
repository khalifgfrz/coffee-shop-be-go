create table public.product (
	product_id serial,
	product_uuid uuid unique default gen_random_uuid(),
	product_name varchar(255) unique not null,
	price int not null,
	category varchar(255) not null,
	description varchar(255),
	stock int not null,
	image varchar,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint product_pk primary key(product_id)
);
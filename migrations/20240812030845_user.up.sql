create table public.user (
	user_id serial,
	user_uuid uuid unique default gen_random_uuid(),
	first_name varchar(255),
	last_name varchar(255),
	phone varchar(255) unique,
	address varchar(255),
	birth_date date,
	image varchar,
	email varchar(255) unique not null,
	password varchar(20) unique not null,
	role varchar(10) not null,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint user_pk primary key(user_id)
);
package models

import "time"

var schemaUser = `
create table public.user (
	user_id serial,
	user_uuid uuid unique default gen_random_uuid(),
	first_name varchar(255),
	last_name varchar(255),
	phone varchar(255) unique,
	address varchar(255),
	birth_date date,
	email varchar(255) unique not null,
	password varchar(20) unique not null,
	role varchar(10) not null,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint user_pk primary key(user_id)
);
`

type User struct {
	User_id	     int     `db:"user_id" json:"user_id"`
	User_uuid	 string     `db:"user_uuid" json:"user_uuid"`
	First_name   string     `db:"first_name" json:"first_name"`
	Last_name    string     `db:"last_name" json:"last_name"`
	Phone		 string     `db:"phone" json:"phone"`
	Address		 string     `db:"address" json:"address"`
	Birth_date 	 string     `db:"birth_date" json:"birth_date"`
	Email	 	 string     `db:"email" json:"email"`
	Password	 string     `db:"password" json:"password"`
	Role	 	 string     `db:"role" json:"role"`
	Created_at   *time.Time `db:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at"`
}

type Users []User

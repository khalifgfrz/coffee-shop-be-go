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
	image varchar,
	email varchar(255) unique not null,
	password varchar(20) unique not null,
	role varchar(10) not null,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint user_pk primary key(user_id)
);
`

type User struct {
	User_id	     string     `db:"user_id" json:"user_id" form:"user_id" valid:"-"`
	User_uuid	 string     `db:"user_uuid" json:"user_uuid" form:"user_uuid" valid:"-"`
	First_name   string     `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    string     `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Phone		 string     `db:"phone" json:"phone" form:"phone" valid:"-"`
	Address		 string     `db:"address" json:"address" form:"address" valid:"stringlength(5|256)~Alamat minimal 5"`
	Birth_date 	 string     `db:"birth_date" json:"birth_date" form:"birth_date" valid:"-"`
	Email	 	 string     `db:"email" json:"email" form:"email" valid:"email"`
	Password	 string     `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role	 	 string     `db:"role" json:"role" form:"role" valid:"-"`
	User_image   string     `db:"user_image" json:"user_image" valid:"-"`
	Created_at   *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type GetUser struct {
	User_id	     string     `db:"user_id" json:"user_id" form:"user_id" valid:"-"`
	User_uuid	 string     `db:"user_uuid" json:"user_uuid" form:"user_uuid" valid:"-"`
	First_name   *string     `db:"first_name" json:"first_name" form:"first_name" valid:"-"`
	Last_name    *string     `db:"last_name" json:"last_name" form:"last_name" valid:"-"`
	Phone		 string     `db:"phone" json:"phone" form:"phone" valid:"-"`
	Address		 *string     `db:"address" json:"address" form:"address" valid:"stringlength(5|256)~Alamat minimal 5"`
	Birth_date 	 *string     `db:"birth_date" json:"birth_date" form:"birth_date" valid:"-"`
	Email	 	 string     `db:"email" json:"email" form:"email" valid:"email"`
	Password	 string     `db:"password" json:"password" form:"password" valid:"stringlength(5|256)~Password minimal 5 karakter"`
	Role	 	 string     `db:"role" json:"role" form:"role" valid:"-"`
	User_image   *string     `db:"user_image" json:"user_image" valid:"-"`
	Created_at   *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type GetUsers []GetUser

type UserQuery struct {
	Page        int    `form:"page"`
}

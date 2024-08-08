package models

import "time"

var schemaProduct = `
create table public.product (
	product_id serial,
	product_uuid uuid unique default gen_random_uuid(),
	product_name varchar(255) unique not null,
	price int not null,
	category varchar(255) not null,
	description varchar(255),
	stock int not null,
	created_at timestamptz default now(),
	updated_at timestamptz,
	constraint product_pk primary key(product_id)
);
`

type Product struct {
	Product_id   string     `db:"product_id" json:"product_id" valid:"-"`
	Product_uuid string     `db:"product_uuid" json:"product_uuid" valid:"-"`
	Product_name string     `db:"product_name" json:"product_name" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Price	     int	    `db:"price" json:"price" valid:"-"`
	Category     string     `db:"category" json:"category" valid:"type(string)"`
	Description  string     `db:"description" json:"description" valid:"type(string)"`
	Stock	     int	    `db:"stock" json:"stock" valid:"-"`
	Created_at   *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type Products []Product

type ProductQuery struct {
	Product_name string `form:"product_name"`
	MinPrice     int    `form:"min_price"`
	MaxPrice     int    `form:"max_price"`
	Category     string `form:"category"`
	SortBy       string `form:"sort_by"`
	Page         int    `form:"page"`
}
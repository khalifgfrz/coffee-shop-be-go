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
	Product_id   int        `db:"product_id" json:"product_id"`
	Product_uuid string     `db:"product_uuid" json:"product_uuid"`
	Product_name string     `db:"product_name" json:"product_name"`
	Price	     int	    `db:"price" json:"price"`
	Category     string     `db:"category" json:"category"`
	Description  string     `db:"description" json:"description"`
	Stock	     int	    `db:"stock" json:"stock"`
	Created_at   *time.Time `db:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at"`
}

type Products []Product

type ProductQuery struct {
	Product_name string `form:"product_name"`
	MinPrice    int    `form:"min_price"`
	MaxPrice    int    `form:"max_price"`
	Category    string `form:"category"`
	SortBy      string `form:"sort_by"`
	Page        int    `form:"page"`
}
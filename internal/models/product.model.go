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
	image varchar,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint product_pk primary key(product_id)
);
`

type Product struct {
	Product_name  string     `db:"product_name" json:"product_name" form:"product_name" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Price	      int	     `db:"price" json:"price" form:"price" valid:"-"`
	Category      string     `db:"category" json:"category" form:"category" valid:"type(string)"`
	Description   string     `db:"description" json:"description" form:"description" valid:"type(string)"`
	Stock	      int	     `db:"stock" json:"stock" form:"stock" valid:"-"`
	Product_image string     `db:"product_image" json:"product_image" valid:"-"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type GetProduct struct {
	Product_id    string     `db:"product_id" json:"product_id" form:"product_id" valid:"-"`
	Product_uuid  string     `db:"product_uuid" json:"product_uuid" form:"product_uuid" valid:"-"`
	Product_name  *string     `db:"product_name" json:"product_name" form:"product_name" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Price	      *int	     `db:"price" json:"price" form:"price" valid:"-"`
	Category      *string     `db:"category" json:"category" form:"category" valid:"type(string)"`
	Description   *string     `db:"description" json:"description" form:"description" valid:"type(string)"`
	Stock	      *int	     `db:"stock" json:"stock" form:"stock" valid:"-"`
	Product_image *string     `db:"product_image" json:"product_image" valid:"-"`
	Created_at    *time.Time `db:"created_at" json:"created_at" valid:"-"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at" valid:"-"`
}

type GetProducts []GetProduct

type ProductQuery struct {
	Product_name string `form:"product_name"`
	MinPrice     int    `form:"min_price"`
	MaxPrice     int    `form:"max_price"`
	Category     string `form:"category"`
	SortBy       string `form:"sort_by"`
	Page         int    `form:"page"`
}
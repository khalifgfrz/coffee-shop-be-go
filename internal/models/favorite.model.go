package models

import "time"

var schemaFavorite = `
create table public.favorite (
	favorite_id serial,
	favorite_uuid uuid unique default gen_random_uuid(),
	user_id int,
	product_id int,
	created_at timestamp without time zone default now(),
	updated_at timestamp without time zone,
	constraint favorite_pk primary key(favorite_id),
	constraint product_fk foreign key (product_id) references public.product(product_id) on delete set null,
	constraint user_fk foreign key (user_id) references public.user(user_id) on delete set null
);
`

type Favorite struct {
	Favorite_id   int        `db:"favorite_id" json:"favorite_id"`
	Favorite_uuid string     `db:"favorite_uuid" json:"favorite_uuid"`
	First_name    string     `db:"first_name" json:"first_name"`
	Last_name     string     `db:"last_name" json:"last_name"`
	Phone         string     `db:"phone" json:"phone"`
	Address       string     `db:"address" json:"address"`
	Email         string     `db:"email" json:"email"`
	Product_name  string     `db:"product_name" json:"product_name"`
	Price	      int	     `db:"price" json:"price"`
	Category      string     `db:"category" json:"category"`
	Description   string     `db:"description" json:"description"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}

type PostFavorite struct {
	Product_id    int        `db:"product_id" json:"product_id"`
	User_id    	  int        `db:"user_id" json:"user_id"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}

type Favorites []Favorite

type UpdateFavorite struct {
	Favorite_id   int        `db:"favorite_id" json:"favorite_id"`
	Favorite_uuid string     `db:"favorite_uuid" json:"favorite_uuid"`
	User_id    	  int        `db:"user_id" json:"user_id"`
	Product_id    int        `db:"product_id" json:"product_id"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}
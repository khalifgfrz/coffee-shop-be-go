package models

import "time"

var schemaOrder = `
CREATE TABLE public.order_list (
	orderlist_id serial4 NOT NULL,
	orderlist_uuid uuid unique DEFAULT gen_random_uuid(),
	user_id int4 NULL,
	subtotal int4 NULL,
	tax int4 NULL,
	payment_id int4 NULL,
	delivery_id int4 NULL,
	status varchar(255) NOT NULL,
	grand_total int4 NULL,
	created_at timestamptz DEFAULT now() NULL,
	updated_at timestamptz NULL,
	CONSTRAINT order_list_uuid_key UNIQUE (orderlist_uuid),
	CONSTRAINT order_pk PRIMARY KEY (orderlist_id),
	CONSTRAINT delivery_fk FOREIGN KEY (delivery_id) REFERENCES public.deliveries(delivery_id) ON DELETE SET null,
	CONSTRAINT payment_fk FOREIGN KEY (payment_id) REFERENCES public.payments(payment_id) ON DELETE SET null,
	CONSTRAINT user_fk FOREIGN KEY (user_id) REFERENCES public.user(user_id) ON DELETE SET null
);
`

type GetOrder struct {
	Orderlist_id    string     `db:"orderlist_id" json:"orderlist_id"`
	Orderlist_uuid  string     `db:"orderlist_uuid" json:"orderlist_uuid"`
	First_name      *string    `db:"first_name" json:"first_name"`
	Last_name       *string    `db:"last_name" json:"last_name"`
	Phone           *string    `db:"phone" json:"phone"`
	Address         *string    `db:"address" json:"address"`
	Email           *string    `db:"email" json:"email"`
	Subtotal	    *string    `db:"subtotal" json:"subtotal"`
	Tax             *string    `db:"tax" json:"tax"`
	Payment_method  *string    `db:"payment_method" json:"payment_method"`
	Delivery_method *string    `db:"deliver_method" json:"deliver_method"`
	Added_cost	    *string    `db:"added_cost" json:"added_cost" valid:"-"`
	Status		    *string    `db:"status" json:"status" valid:"-"`
	Grand_total	    *string    `db:"grand_total" json:"grand_total" valid:"-"`
	Created_at      *time.Time `db:"created_at" json:"created_at"`
	Updated_at      *time.Time `db:"updated_at" json:"updated_at"`
}

type Order struct {
	User_id    	  int        `db:"user_id" json:"user_id"`
	Subtotal      int        `db:"subtotal" json:"subtotal"`
	Tax		      int        `db:"tax" json:"tax"`
	Payment_id 	  int        `db:"payment_id" json:"payment_id"`
	Delivery_id	  int        `db:"delivery_id" json:"delivery_id"`
	Status		  string     `db:"status" json:"status"`
	Grand_total	  int        `db:"grand_total" json:"grand_total"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}

type GetOrders []GetOrder

type OrderQuery struct {
	Page        int    `form:"page"`
}
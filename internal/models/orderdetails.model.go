package models

var schemaOrderDetails = `
CREATE TABLE public.order_details (
	orderdetails_id serial4 NOT NULL,
	size_id int4 NULL,
	order_id int4 NULL,
	product_id int4 NULL,
	qty int4 NULL,
	CONSTRAINT orderdetails_pk PRIMARY KEY (orderdetails_id),
	CONSTRAINT order_fk FOREIGN KEY (order_id) REFERENCES public.order_list(orderlist_id) ON DELETE SET null,
	CONSTRAINT product_fk FOREIGN KEY (product_id) REFERENCES public.product(product_id) ON DELETE SET null,
	CONSTRAINT size_fk FOREIGN KEY (size_id) REFERENCES public.sizes(size_id) ON DELETE SET NULL
);
`

type ProductDetail struct {
	Size_id    int `db:"size_id" json:"size_id"`
	Product_id int `db:"product_id" json:"product_id"`
	Qty        int `db:"qty" json:"qty"`
}

type GetOrderDetail struct {
	Order_id      *int    `db:"order_id" json:"order_id"`
	Product_image *string `db:"product_image" json:"product_image" valid:"-"`
	Product_name  *string `db:"product_name" json:"product_name" form:"product_name" valid:"stringlength(5|100)~Nama Product minimal 5 dan maksimal 100"`
	Price         *int    `db:"price" json:"price" form:"price" valid:"-"`
	Size          *string `db:"size" json:"size" form:"size" valid:"-"`
	Added_cost    *int    `db:"added_cost" json:"added_cost" form:"added_cost" valid:"-"`
	Qty           *int    `db:"qty" json:"qty"`
}

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
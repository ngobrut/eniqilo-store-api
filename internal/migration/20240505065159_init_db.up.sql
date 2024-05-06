CREATE TABLE IF NOT EXISTS users (
    user_id uuid default gen_random_uuid() not null constraint users_pk primary key,
    name varchar(50) not null,
    phone varchar(25) not null,
    password varchar(500) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
CREATE UNIQUE INDEX unique_user_phone_idx ON users (phone);
    
CREATE TABLE IF NOT EXISTS customers (
    customer_id uuid default gen_random_uuid() not null constraint customers_pk primary key,
    name varchar(50) not null,
    phone varchar(25) not null,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
CREATE UNIQUE INDEX unique_customer_phone_idx on customers (phone);

CREATE TABLE IF NOT EXISTS products (
	product_id uuid default gen_random_uuid() not null constraint products_pk primary key,
	name varchar(30) not null,
	sku varchar(30) not null,
	category  varchar(30) not null,
	image_url text not null,
	notes varchar(200) not null,
	price int not null,
	stock int not null,
	location varchar(200),
	is_available bool not null,
	created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
);

CREATE TABLE IF NOT EXISTS checkouts (
	checkout_id uuid default gen_random_uuid() not null constraint checkouts_pk primary key,
	customer_id uuid not null,
	total_price int not null,
	paid int not null,
	change int not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	constraint customer_id_fk foreign key (customer_id) references customers(customer_id)
);

CREATE TABLE IF NOT EXISTS checkout_products (
	id serial not null constraint checkout_product_id_pk primary key,
	checkout_id uuid not null,
	product_id uuid not null,
	quantity int not null,
	price int not null,
	created_at timestamp default now(),
	constraint cp_checkout_id_fk foreign key (checkout_id) references checkouts(checkout_id),
	constraint cp_product_id_fk foreign key (product_id) references products(product_id)
);
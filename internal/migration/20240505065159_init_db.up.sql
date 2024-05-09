CREATE TABLE IF NOT EXISTS users (
	user_id uuid default gen_random_uuid() not null constraint users_pk primary key,
	name varchar(50) not null,
	phone varchar(25) not null,
	password varchar(500) not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	deleted_at timestamp default null
);

CREATE UNIQUE INDEX unique_user_phone_idx ON users (phone)
WHERE deleted_at IS NULL;

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
	category varchar(30) not null,
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
CREATE INDEX idx_sku ON products(sku);
CREATE INDEX idx_name ON products(name);

CREATE TABLE IF NOT EXISTS invoices (
	invoice_id uuid default gen_random_uuid() not null constraint invoice_pk primary key,
	customer_id uuid not null,
	total_price int not null,
	paid int not null,
	change int not null,
	created_at timestamp default now(),
	updated_at timestamp default now(),
	constraint customer_id_fk foreign key (customer_id) references customers(customer_id)
);

CREATE INDEX idx_invoices_customer_id ON invoices(customer_id);

CREATE TABLE IF NOT EXISTS invoice_products (
	id serial not null constraint invoice_product_id_pk primary key,
	invoice_id uuid not null,
	product_id uuid not null,
	quantity int not null,
	price int not null,
	created_at timestamp default now(),
	constraint ip_invoice_id_fk foreign key (invoice_id) references invoices(invoice_id),
	constraint ip_product_id_fk foreign key (product_id) references products(product_id)
);

CREATE INDEX idx_invoice_products_invoice_id ON invoice_products(invoice_id);
CREATE INDEX idx_invoice_products_product_id ON invoice_products(product_id);

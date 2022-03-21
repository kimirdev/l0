CREATE TABLE deliveries
(
	id serial not null unique,
	name varchar(255) NOT NULL,
	phone varchar(255) NOT NULL,
	zip varchar(255) NOT NULL,
	city varchar(255) NOT NULL,
	address varchar(255) NOT NULL,
	region varchar(255) NOT NULL,
	email varchar(255) NOT NULL
);

CREATE TABLE payments
(
	id serial NOT NULL UNIQUE,
	transaction varchar(255) NOT NULL,
	request_id varchar(255),
	currency varchar(255) NOT NULL,
	provider varchar(255) NOT NULL,
	amount bigint NOT NULL,
	payment_dt bigint NOT NULL,
	bank varchar(255) NOT NULL,
	delivery_cost bigint NOT NULL,
	goods_total bigint NOT NULL,
	custom_fee bigint NOT NULL
);

CREATE TABLE items
(
	id serial NOT NULL UNIQUE,
	chrt_id bigint NOT NULL,
	track_number varchar(255) NOT NULL,
	price bigint NOT NULL,
	rid varchar(255) NOT NULL,
	name varchar(255) NOT NULL,
	sale bigint NOT NULL,
	size varchar NOT NULL,
	total_price bigint NOT NULL,
	nm_id bigint NOT NULL,
	brand varchar(255) NOT NULL,
	status bigint NOT NULL
);

CREATE TABLE orders
(
	id serial NOT NULL UNIQUE,
	order_uid varchar(255) NOT NULL UNIQUE,
	track_number varchar(255) NOT NULL,
	entry varchar(255) NOT NULL,
	local varchar(255) NOT NULL,
	internal_signature varchar(255) NOT NULL,
	customer_id varchar(255) NOT NULL,
	delivery_service varchar(255) NOT NULL,
	shardkey varchar(255) NOT NULL,
	sm_id int NOT NULL,
	date_created varchar(255) NOT NULL,
	oof_shard varchar(255) NOT NULL,

	delivery_id int references deliveries(id) on delete cascade NOT NULL,
	payment_id int references payments(id) on delete cascade NOT NULL

);

CREATE TABLE items_orders
(
	id serial not null unique,
	order_id int references orders(id) on delete cascade not null,
	item_id int references items(id) on delete cascade not null
);
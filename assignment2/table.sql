create table items (
    item_id int not null,
    item_code varchar(255) not null,
    description varchar(255),
    quantity int,
    order_id int,
    primary key (item_id),
    foregign key (order_id) references orders(order_id)
);

create table orders (
    order_id int not null,
    customer_name varchar(255),
    ordered_at date,
    primary key(order_id)
);
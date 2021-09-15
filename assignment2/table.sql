create table orders_by (
    items varchar(255) not null,
    item_id int not null,
    item_code varchar(255) not null,
    description varchar(255),
    quantity int,
    order_id int,
    primary key (order_id)
)

create table orders (
    orders varchar(255) not null,
    order_id int not null,
    customer_name varchar(255),
    ordered_at date,
    
)
create schema schema_example;

create table schema_example.products
(
    id         bigint not null
        constraint products_pk
            primary key,
    name       varchar(50),
    price      numeric,
    created_at date default now()
);

INSERT INTO schema_example.products (id, name, price, created_at) VALUES (1, 'tv', 199.9, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (2, 'radio', 49.9, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (3, 'pc', 699, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (4, 'smartphone', 299.99, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (5, 'tablet', 199, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (6, 'headphones', 99.99, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (7, 'digital camera', 159, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (8, 'printer', 129.99, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (9, 'smart watch', 249, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (10, 'video games console', 399, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (11, 'wireless keyboard', 49.99, '2024-07-17');
INSERT INTO schema_example.products (id, name, price, created_at) VALUES (12, 'mouse', 19.99, '2024-07-17');

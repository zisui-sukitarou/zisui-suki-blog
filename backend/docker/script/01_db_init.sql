CREATE DATABASE IF NOT EXISTS zisuidb;
USE zisuidb;
/*
CREATE DATABASE zisuidb;
\c zisuidb;
CREATE SCHEMA zisuischema;
CREATE ROLE zisui WITH LOGIN PASSWORD 'password';
GRANT ALL PRIVILEGES ON SCHEMA zisuischema TO zisui;
*/


/*
DROP TABLE IF EXISTS customers;
CREATE TABLE customers(
    customer_id varchar(32) NOT NULL,
    name varchar(32) NOT NULL,
    PRIMARY KEY (customer_id)
);

DROP TABLE IF EXISTS products;
CREATE TABLE products(
    product_id varchar(32) NOT NULL,
    name varchar(32) NOT NULL,
    price int NOT NULL,
    PRIMARY KEY (product_id)
);

DROP TABLE IF EXISTS customer_products;
CREATE TABLE customer_products(
    customer_id varchar(32) NOT NULL,
    product_id varchar(32) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
    FOREIGN KEY (product_id) REFERENCES products(product_id),
    PRIMARY KEY (customer_id, product_id)
);
*/
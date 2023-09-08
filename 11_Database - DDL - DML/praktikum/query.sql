-- Active: 1693796158705@@127.0.0.1@3306
CREATE DATABASE alta_online_shop
    DEFAULT CHARACTER SET = 'utf8mb4';

SHOW DATABASES;

use alta_online_shop;

CREATE TABLE operators(
    id VARCHAR(5) NOT NULL, 
    name VARCHAR(50) NOT NULL, 
    username VARCHAR(50) NOT NULL, 
    password VARCHAR(50) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

SHOW CREATE TABLE operators;

CREATE TABLE product_type(
    id VARCHAR(5) NOT NULL,
    category VARCHAR(20) NOT NULL,
    PRIMARY KEY(id)
);

SHOW CREATE TABLE product_type;

CREATE TABLE product_descriptions(
    id VARCHAR(5) not NULL,
    description VARCHAR(100),
    PRIMARY KEY(id)
);

SHOW CREATE TABLE product_descriptions;

CREATE TABLE products(
    id VARCHAR(5) NOT NULL,
    name VARCHAR(30) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    input_by VARCHAR(5) NOT NULL,
    category VARCHAR(5) NOT NULL,
    description VARCHAR(5),
    PRIMARY key(id)
);

SHOW CREATE TABLE products;

CREATE TABLE users(
    id VARCHAR(5) NOT NULL,
    name VARCHAR(50) NOT NULL,
    address VARCHAR(50) NOT NULL,
    date_of_birth DATE NOT NULL,
    status ENUM('single', 'married') NOT NULL,
    gender ENUM('male', 'female') NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

SHOW CREATE TABLE users;

CREATE Table payment_method(
    id VARCHAR(5) NOT NULL,
    name VARCHAR(20) NOT NULL,
    PRIMARY KEY(id)
);

SHOW CREATE TABLE payment_method;

CREATE TABLE transaction(
    id VARCHAR(5) NOT NULL PRIMARY KEY,
    id_user VARCHAR(5) NOT NULL,
    id_payment VARCHAR(5) NOT NULL,
    total INT NOT NULL DEFAULT 0,
    date DATETIME DEFAULT CURRENT_TIMESTAMP
);

show tables;
SHOW CREATE TABLE transaction;

CREATE TABLE transaction_details(
    id_transaction VARCHAR(5) NOT NULL,
    id_product VARCHAR(5) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    subtotal INT NOT NULL default 0
);

SHOW CREATE TABLE details_transaction;


CREATE TABLE kurir(
    id VARCHAR(5) NOT NULL PRIMARY KEY,
    name VARCHAR(50),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

SHOW CREATE TABLE kurir;
DESC kurir;

ALTER TABLE kurir ADD COLUMN ongkos_dasar INT NOT NULL DEFAULT 0;

ALTER TABLE kurir RENAME shipping;

SHOW TABLES;

DROP TABLE shipping;

DESC products;

SHOW CREATE TABLE product_descriptions;

SHOW CREATE TABLE products;

ALTER TABLE product_descriptions ADD
    CONSTRAINT fk_product_id
        FOREIGN KEY(id) REFERENCES products(id);

ALTER TABLE products ADD
    CONSTRAINT fk_input_by
        FOREIGN KEY(input_by) REFERENCES operators(id);

ALTER TABLE products ADD
    CONSTRAINT fk_product_type
        FOREIGN KEY(category) REFERENCES product_type(id);

DESC transaction_details;

ALTER TABLE transaction_details ADD
    PRIMARY KEY(id_transaction, id_product);

ALTER TABLE transaction_details ADD
    CONSTRAINT fk_id_transaction 
        FOREIGN KEY(id_transaction) REFERENCES transaction(id);

ALTER TABLE transaction_details ADD
    CONSTRAINT fk_id_product
        FOREIGN KEY(id_product) REFERENCES products(id);

SHOW CREATE TABLE transaction_details;

ALTER TABLE transaction ADD
    CONSTRAINT fk_id_user 
        FOREIGN KEY(id_user) REFERENCES users(id);

ALTER TABLE transaction ADD 
    CONSTRAINT fk_id_payment
        FOREIGN KEY(id_payment) REFERENCES payment_method(id);


SHOW CREATE TABLE transaction;
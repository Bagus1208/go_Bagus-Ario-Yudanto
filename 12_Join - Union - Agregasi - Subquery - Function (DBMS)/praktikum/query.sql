-- Active: 1693796158705@@127.0.0.1@3306@alta_online_shop
SHOW DATABASES;

CREATE DATABASE olshop;

USE olshop;

CREATE TABLE product_types(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE operators(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE product_descriptions(
    id INT(11) NOT NULL AUTO_INCREMENT,
    description TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE products(
    id INT(11) NOT NULL AUTO_INCREMENT,
    product_type_id INT(11) NOT NULL,
    operator_id INT(11) NOT NULL,
    code VARCHAR(50),
    name VARCHAR(100) NOT NULL,
    status SMALLINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

ALTER TABLE product_descriptions ADD
    CONSTRAINT fk_id
        FOREIGN KEY(id) REFERENCES products(id);

ALTER TABLE products ADD
    CONSTRAINT fk_product_type_id
        FOREIGN KEY(product_type_id) REFERENCES product_type(id);

ALTER TABLE products ADD
    CONSTRAINT fk_operator_id
        FOREIGN KEY(operator_id) REFERENCES operators(id);

CREATE TABLE payment_methods(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE users(
    id INT(11) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    dob DATE NOT NULL,
    gender CHAR(1),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

CREATE TABLE transactions(
    id INT(11) NOT NULL AUTO_INCREMENT,
    user_id INT(11) NOT NULL,
    payment_method_id INT(11) NOT NULL,
    status VARCHAR(10) NOT NULL,
    total_qty INT(11) NOT NULL DEFAULT 0,
    total_price NUMERIC(25,2) NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
ALTER TABLE transactions ADD
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id) REFERENCES users(id);

ALTER TABLE transactions ADD
    CONSTRAINT fk_payment_method_id
        FOREIGN KEY(payment_method_id) REFERENCES payment_methods(id);

CREATE TABLE transaction_details(
    transaction_id INT(11) NOT NULL,
    product_id INT(11) NOT NULL,
    status VARCHAR(10) NOT NULL,
    qty INT(11) NOT NULL DEFAULT 0,
    price NUMERIC(25,2) NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(transaction_id, product_id)
);

ALTER TABLE transaction_details ADD
    CONSTRAINT fk_transaction_id
        FOREIGN KEY(transaction_id) REFERENCES transactions(id);

ALTER TABLE transaction_details ADD
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id) REFERENCES products(id);

INSERT INTO operators(name) VALUES 
    ('Rendy'),
    ('Okta'),
    ('Budi'),
    ('Sandy'),
    ('Putri');

SELECT * FROM operators;

INSERT INTO product_types(name) VALUES
    ('Food'),
    ('Drink'),
    ('Snack');

SELECT * FROM product_types;

INSERT INTO products(product_type_id, operator_id, name, status) VALUES
    (1, 3, 'Fried Chicken', 100),
    (1, 3, 'Meatball', 100);


INSERT INTO products(product_type_id, operator_id, name, status) VALUES
    (2, 1, 'Ice Tea', 200),
    (2, 1, 'Cola', 200),
    (2, 1, 'Mineral water', 1000);

INSERT INTO products(product_type_id, operator_id, name, status) VALUES
    (3, 4, 'Chips', 400),
    (3, 4, 'French fries', 250),
    (3, 4, 'Pancake', 150);

SELECT * FROM products;

INSERT INTO product_descriptions(description) VALUES
    ('chicken meat dipped in batter and fried until the skin becomes crispy and the meat is tender'),
    ('beef that is processed and shaped into balls served with soup'),
    ('sweet tea is served cold'),
    ('carbonated soft drink served chilled and sweet and refreshing taste.'),
    ('natural spring water that contains a variety of minerals'),
    ('the potatoes are thinly sliced ​​and then fried so they have a crunchy texture and then have a savory taste.'),
    ('potatoes that are cut and then fried have a crunchy texture on the outside and soft on the inside'),
    ('thin flat cake served with a topping of syrup, butter or fruit.');

SELECT * FROM product_descriptions;

INSERT INTO payment_methods(name, status) VALUES
    ('Transfer', 1),
    ('QRIS', 1),
    ('E-Wallet', 1);

SELECT * FROM payment_methods;

INSERT INTO users(name, status, dob, gender) VALUES
    ('Fahri', 1,  '1990-07-22', 'M'),
    ('Ratu', 1, '1998-11-01', 'F'),
    ('Rio', 1, '2001-07-17', 'M'),
    ('Anya', 1, '1999-03-04', 'F'),
    ('Bagus', 1, '2002-08-12', 'M');

SELECT * FROM users;

INSERT INTO transactions(user_id, payment_method_id, status, total_qty, total_price) VALUES
    (1, 1, 'success', 3, 25000),
    (1, 2, 'success', 3, 30000),
    (1, 3, 'success', 3, 50000),
    (2, 1, 'success', 4, 50000),
    (2, 2, 'success', 5, 50000),
    (2, 3, 'success', 7, 80000),
    (3, 1, 'success', 3, 30000),
    (3, 2, 'success', 6, 60000),
    (3, 3, 'success', 7, 70000),
    (4, 1, 'success', 3, 25000),
    (4, 2, 'success', 8, 90000),
    (4, 3, 'success', 3, 30000),
    (5, 1, 'success', 5, 60000),
    (5, 2, 'success', 3, 40000),
    (5, 3, 'success', 4, 40000);

SELECT * FROM transactions;

INSERT INTO transaction_details(transaction_id, product_id, status, qty, price) VALUES
    (1, 1, 'success', 1, 15000),
    (1, 3, 'success', 1, 5000),
    (1, 6, 'success', 1, 5000),
    (2, 1, 'success', 1, 15000),
    (2, 3, 'success', 1, 5000),
    (2, 7, 'success', 1, 10000),
    (3, 2, 'success', 1, 20000),
    (3, 4, 'success', 1, 5000),
    (3, 6, 'success', 1, 5000),
    (4, 1, 'success', 2, 30000),
    (4, 3, 'success', 1, 5000),
    (4, 8, 'success', 1, 15000),
    (5, 2, 'success', 1, 20000),
    (5, 3, 'success', 2, 10000),
    (5, 7, 'success', 2, 20000),
    (6, 2, 'success', 3, 60000),
    (6, 3, 'success', 3, 15000),
    (6, 6, 'success', 1, 5000),
    (7, 2, 'success', 1, 20000),
    (7, 4, 'success', 1, 5000),
    (7, 6, 'success', 1, 5000),
    (8, 1, 'success', 3, 45000),
    (8, 5, 'success', 2, 10000),
    (8, 6, 'success', 1, 5000),
    (9, 1, 'success', 3, 45000),
    (9, 6, 'success', 3, 15000),
    (9, 7, 'success', 1, 10000),
    (10, 1, 'success', 1, 15000),
    (10, 3, 'success', 1, 5000),
    (10, 7, 'success', 1, 10000),
    (11, 2, 'success', 2, 40000),
    (11, 4, 'success', 2, 10000),
    (11, 7, 'success', 4, 40000),
    (12, 2, 'success', 1, 20000),
    (12, 4, 'success', 1, 5000),
    (12, 6, 'success', 1, 5000),
    (13, 1, 'success', 3, 45000),
    (13, 5, 'success', 1, 5000),
    (13, 7, 'success', 1, 10000),
    (14, 2, 'success', 1, 20000),
    (14, 3, 'success', 1, 5000),
    (14, 8, 'success', 1, 15000),
    (15, 1, 'success', 1, 15000),
    (15, 3, 'success', 1, 5000),
    (15, 7, 'success', 2, 20000);

SELECT * FROM transaction_details;

SELECT name FROM users where gender = 'M';

SELECT * FROM products WHERE id = 3;

SELECT * FROM users 
WHERE created_at >= DATE_ADD(CURDATE(), INTERVAL -7 DAY)
AND name LIKE '%a%';

SELECT COUNT(gender) AS count_gender_female FROM users
WHERE gender = 'F';

SELECT * FROM users
ORDER BY name ASC;

SELECT * FROM products
limit 5;


UPDATE products SET name = 'product dummy' WHERE id = 1;
SELECT * FROM products WHERE id = 1;

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;
SELECT * FROM transaction_details WHERE product_id = 1;


SHOW CREATE TABLE transaction_details;
SHOW CREATE TABLE product_descriptions;

ALTER TABLE transaction_details DROP
    CONSTRAINT fk_product_id;

ALTER TABLE product_descriptions DROP
    CONSTRAINT fk_product_id;

ALTER TABLE transaction_details ADD
    CONSTRAINT fk_product_id
        FOREIGN KEY(product_id) REFERENCES products(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE product_descriptions ADD
    CONSTRAINT fk_id
        FOREIGN KEY(id) REFERENCES products(id)
            ON DELETE CASCADE ON UPDATE CASCADE;

DELETE FROM products WHERE id = 1;

SELECT * FROM products;

DELETE FROM products WHERE product_type_id = 1;


SELECT * FROM transactions 
JOIN users ON (users.id = transactions.user_id)
WHERE users.id IN(1,2);

SELECT SUM(total_price) AS total_price 
FROM transactions
WHERE user_id = 1;

SELECT COUNT(transactions.id) AS total_transaction_product_type_2
FROM transactions
JOIN transaction_details AS t ON(t.transaction_id = transactions.id)
JOIN products AS p ON(p.id = t.product_id)
WHERE product_type_id = 2;

SELECT products.*, product_types.name 
FROM products
JOIN product_types ON(product_types.id = products.product_type_id);

SELECT transactions.*, products.name, users.name
FROM transactions
LEFT JOIN transaction_details AS td ON(td.transaction_id = transactions.id) 
LEFT JOIN products ON(products.id = td.product_id)
JOIN users ON(users.id = transactions.user_id);

SELECT * FROM transactions;
SELECT * FROM transaction_details;
SHOW CREATE TABLE transaction_details;

DELIMITER $$
CREATE TRIGGER delete_data_transaction_details
BEFORE DELETE ON transactions FOR EACH ROW
BEGIN
DECLARE data_id INT;
SET data_id = OLD.id;
    DELETE FROM transaction_details WHERE transaction_id = data_id;
END $$
DELIMITER;

SHOW TRIGGERS;

DELIMITER $$
CREATE TRIGGER update_total_qty_transaction
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
UPDATE transactions SET total_qty = total_qty - OLD.qty WHERE id = OLD.transaction_id;
END $$
DELIMITER;

SELECT * FROM products 
WHERE id NOT IN 
    (SELECT product_id 
        FROM transaction_details 
        GROUP BY product_id);

INSERT INTO products(product_type_id, operator_id, name, status) VALUES
(1, 3, 'Ramen', 100),
(3, 2, 'Takoyaki', 200);


SELECT product_id FROM transaction_details ;

SELECT * FROM products;

SELECT * FROM transaction_details;
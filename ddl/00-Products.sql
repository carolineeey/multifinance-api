CREATE TABLE IF NOT EXISTS products
(
    product_id     VARCHAR(255)    NOT NULL PRIMARY KEY,
    product_name   VARCHAR(255)    NOt NULL,
    cost           DECIMAL(15, 2)  NOT NULL
)

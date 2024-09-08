CREATE TABLE IF NOT EXISTS transactions
(
    nik                 INT             NOT NULL,
    transaction_id      VARCHAR(255)    NOT NULL,
    product_id          VARCHAR(255)    NOT NULL,
    admin_fee           DECIMAL(15, 2)  NOT NULL,
    installment_amount  DECIMAL(15, 2)  NOT NULL,
    interest_amount     DECIMAL(4, 2)   NOT NULL,
    created_at          DATETIME        NOT NULL,
    PRIMARY KEY (nik, transaction_id),
    FOREIGN KEY (nik) REFERENCES customers (nik) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (product_id)
)
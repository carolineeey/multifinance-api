CREATE TABLE IF NOT EXISTS customer_limits
(
    nik         INT             NOT NULL PRIMARY KEY,
    loan_term   INT             NOT NULL,
    loan_limit  DECIMAL(15, 2)  NOT NULL,
    FOREIGN KEY (nik) REFERENCES customers (nik) ON DELETE CASCADE ON UPDATE CASCADE
)

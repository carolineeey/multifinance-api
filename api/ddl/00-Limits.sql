CREATE TABLE IF NOT EXISTS limits
(
    salary      DECIMAL(15, 2)  NOT NULL PRIMARY KEY,
    loan_term   INT             NOT NULL,
    limit       DECIMAL(15, 2)  NOT NULL
)

CREATE TABLE IF NOT EXISTS customers
(
    nik          INT             NOT NULL PRIMARY KEY,
    full_name    VARCHAR(255)    DEFAULT NULL,
    legal_name   VARCHAR(255)    DEFAULT NULL,
    birth_place  VARCHAR(255)    DEFAULT NULL,
    birth_date   DATE            DEFAULT NULL,
    salary       DECIMAL(15, 2)  NOT NULL,
    ktp_photo_id VARCHAR(255)    DEFAULT NULL,
    selfie_id    VARCHAR(255)    DEFAULT NULL
)

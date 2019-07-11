CREATE TABLE coupons (
    id SERIAL NOT NULL,
    name varchar(255) NOT NULL,
    brand varchar(255) NOT NULL,
    value real NOT NULL,
    createdAt timestamp NOT NULL,
    expiry timestamp NOT NULL,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX UK_cupons_name ON coupons (name);
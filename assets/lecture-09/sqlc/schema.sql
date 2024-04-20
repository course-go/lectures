CREATE TABLE product (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  price  numeric DEFAULT 0.00
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users  (
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    phone      VARCHAR(20)  NOT NULL,
    track_id   VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS products  (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS orders  (
    id          BIGSERIAL PRIMARY KEY,
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount      NUMERIC(15,2) NOT NULL CHECK (amount >= 0),
    description TEXT,
    products_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_items  (
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id UUID   NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity   INT    NOT NULL CHECK (quantity > 0),
    UNIQUE(order_id, product_id)
);

CREATE INDEX idx_orders_user_id ON orders(user_id);

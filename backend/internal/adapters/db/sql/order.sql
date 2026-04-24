-- name: CreateOrder :one
INSERT INTO orders (user_id, amount, description)
VALUES ($1, $2, $3)
RETURNING id, user_id, amount, description;

-- name: GetOrder :one
SELECT id, user_id, amount, description
FROM orders
WHERE id = $1;

-- name: GetUserOrderProducts :many
SELECT p.id, p.name, p.description
FROM order_items oi
JOIN products p ON p.id = oi.product_id
WHERE oi.order_id = $1;

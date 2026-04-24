-- name: NewProduct :exec
INSERT INTO products (id, name, description)
VALUES ($1, $2, $3);

-- name: EditProduct :exec
UPDATE products
SET name = $2, description = $3
WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;

-- name: GetProduct :one
SELECT id, name, description
FROM products
WHERE id = $1;

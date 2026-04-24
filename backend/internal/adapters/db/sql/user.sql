-- name: SaveUser :exec
INSERT INTO users (id, name, email, phone, track_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (id) DO UPDATE SET
    name       = EXCLUDED.name,
    email      = EXCLUDED.email,
    phone      = EXCLUDED.phone,
    track_id   = EXCLUDED.track_id;

-- name: GetAllUsers :many
SELECT id, name, email, phone, track_id
FROM users
ORDER BY name;

-- name: GetUserOrders :many
SELECT id, user_id, amount, description
FROM orders
WHERE user_id = $1
ORDER BY id DESC;

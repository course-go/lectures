-- name: GetByID :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: GetAll :many
SELECT * FROM product
ORDER BY name;

-- name: Create :one
INSERT INTO product (
  name, price
) VALUES ($1, $2)
RETURNING *;

-- name: Update :exec
UPDATE product
  set name = $2,
  price = $3
WHERE id = $1;

-- name: Delete :exec
DELETE FROM product
WHERE id = $1;

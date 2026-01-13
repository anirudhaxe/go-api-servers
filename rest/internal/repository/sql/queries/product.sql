-- name: GetAllProducts :many
SELECT
    *
FROM
    products;

-- name: CreateProduct :exec
INSERT INTO products (
    id,
    name,
    description,
    price)
VALUES (
    $1,
    $2,
    $3,
    $4);


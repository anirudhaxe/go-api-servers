-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    encrypted_password,
    role,
    is_active)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6)
RETURNING
    *;

-- name: GetUser :one
SELECT
    id,
    username,
    email,
    encrypted_password,
    ROLE,
    is_active,
    created_at,
    last_login_at
FROM
    users
WHERE
    email = $1;


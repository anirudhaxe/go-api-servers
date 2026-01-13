-- name: CreateUser :one
INSERT INTO users (
    id,
    name)
VALUES (
    $1,
    $2)
RETURNING
    *;

-- name: CreateTodo :one
INSERT INTO todos (
    id,
    user_id,
    text,
    done)
VALUES (
    $1,
    $2,
    $3,
    $4)
RETURNING
    *;

-- name: GetTodos :many
SELECT
    *
FROM
    todos;

-- name: GetUsers :many
SELECT
    *
FROM
    users;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = $1;


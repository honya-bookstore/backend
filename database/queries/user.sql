-- name: CreateUser :one
INSERT INTO users (
  id
) VALUES (
  sql.arg('id')
)
RETURNING
  *;

-- name: ListUsers :many
SELECT
  *
FROM
  users
ORDER BY
  id ASC;

-- name: GetUserByID :one
SELECT
  *
FROM
  users
WHERE
  id = sql.arg('id');

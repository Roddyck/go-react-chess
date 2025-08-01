-- name: CreateUser :one
INSERT INTO users (
    id, created_at, updated_at, name, email, hashed_password
) VALUES ( gen_random_uuid(), now(), now(), $1, $2, $3 )
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

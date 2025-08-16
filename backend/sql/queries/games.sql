-- name: CreateGame :one
INSERT INTO games (
    id, board, turn, history, players
) values (
    gen_random_uuid(), $1, $2, $3, $4
)
RETURNING *;

-- name: GetGameByID :one
SELECT * FROM games
WHERE id = $1;

-- name: UpdateGame :exec
UPDATE games
    SET board = $2,
    turn = $3,
    history = $4,
    players = $5
WHERE id = $1;

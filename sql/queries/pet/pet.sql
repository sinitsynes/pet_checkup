-- name: Create :one
INSERT INTO pet (name, birth, passport, chip)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetByID :one
SELECT * FROM pet WHERE id = $1;

-- name: UpdateByID :one
UPDATE pet SET
    name = coalesce(sqlc.narg('name'), name),
    birth = coalesce(sqlc.narg('birth'), birth),
    passport = coalesce(sqlc.narg('passport'), passport),
    chip = coalesce(sqlc.narg('chip'), chip)
WHERE id = @id
RETURNING *;

-- name: DeleteByID :exec
DELETE FROM pet WHERE id = $1;

-- name: GetMany :many
SELECT * FROM pet;

-- name: Create :one
INSERT INTO pets (name, birth, passport, chip)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetByID :one
SELECT * FROM pets WHERE id = $1;

-- name: UpdateByID :one
UPDATE pets SET
    name = coalesce(sqlc.narg('name'), name),
    birth = coalesce(sqlc.narg('birth'), birth),
    passport = coalesce(sqlc.narg('passport'), passport),
    chip = coalesce(sqlc.narg('chip'), chip)
WHERE id = @id
RETURNING *;

-- name: DeleteByID :exec
DELETE FROM pets WHERE id = $1;

-- name: GetAll :many
SELECT * FROM pets;

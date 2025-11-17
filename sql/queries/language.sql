-- name: Create :one
INSERT INTO languages (name, label)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteByID :exec
DELETE FROM languages
WHERE id = @id;

-- name: GetAll :many
SELECT *
FROM languages;

-- name: GetByID :one
SELECT *
FROM languages
WHERE id = @id;

-- name: UpdateByID :one
UPDATE languages
SET name = coalesce(sqlc.narg('name'), name),
    label = coalesce(sqlc.narg('label'), label)
WHERE id = @id
RETURNING *;

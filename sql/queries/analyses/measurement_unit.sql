-- name: Create :one
INSERT INTO measurement_unit (name)
VALUES ($1)
RETURNING *;

-- name: UpdateByID :one
UPDATE measurement_unit
SET name = @name
WHERE id = @id
RETURNING *;

-- name: GetMany :many
SELECT id, name
FROM measurement_unit;

-- name: GetByID :one
SELECT id, name
FROM measurement_unit
WHERE id = @id;

-- name: DeleteByID :exec
DELETE FROM measurement_unit
WHERE id = @id;

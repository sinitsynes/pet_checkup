-- name: Create :one
INSERT INTO panel_measurement_units (name)
VALUES ($1)
RETURNING *;

-- name: UpdateByID :one
UPDATE panel_measurement_units
SET name = @name
WHERE id = @id
RETURNING *;

-- name: GetMany :many
SELECT id, name
FROM panel_measurement_units;

-- name: GetByID :one
SELECT id, name
FROM panel_measurement_units
WHERE panel_measurement_units.id = @id;

-- name: DeleteByID :exec
DELETE FROM panel_measurement_units
WHERE panel_measurement_units.id = @id;

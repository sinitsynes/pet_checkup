-- name: Create :one
INSERT INTO component (name, unit_id, min_amount, max_amount, comment)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdateByID :one
UPDATE component
SET name = coalesce(sqlc.narg('name'), name),
    unit_id = coalesce(sqlc.narg('unit_id'), unit_id),
    min_amount = coalesce(sqlc.narg('min_amount'), min_amount),
    max_amount = coalesce(sqlc.narg('max_amount'), max_amount),
    comment = coalesce(sqlc.narg('comment'), comment)
WHERE id = $1
RETURNING *;

-- name: GetByID :one
SELECT component.id,
	   component.name,
	   json_build_object('id', component.unit_id, 'name', measurement_unit.name) AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM component
JOIN measurement_unit ON component.unit_id = measurement_unit.id
WHERE component.id = $1;

-- name: GetMany :many
SELECT component.id,
	   component.name,
	   json_build_object('id', component.unit_id, 'name', measurement_unit.name) AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM component
JOIN measurement_unit ON component.unit_id = measurement_unit.id;

-- name: DeleteByID :exec
DELETE FROM component
WHERE id = $1;

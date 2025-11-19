-- name: Create :one
INSERT INTO component (name, unit_id, min_amount, max_amount, comment)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetByID :one
SELECT component.id,
	   component.name,
	   measurement_unit.name AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM component
JOIN measurement_unit ON component.unit_id = measurement_unit.id
WHERE component.id = $1;

-- name: GetMany :many
SELECT component.id,
	   component.name,
	   measurement_unit.name AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM component
JOIN measurement_unit ON component.unit_id = measurement_unit.id;

-- name: Create :one
INSERT INTO panel_components (name, unit_id, min_amount, max_amount, comment)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetOne :one
SELECT panel_components.id,
	   panel_components.name,
	   panel_measurement_units.name AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM panel_components
JOIN panel_measurement_units ON panel_components.unit_id = panel_measurement_units.id
WHERE panel_components.id = $1;

-- name: GetMany :many
SELECT panel_components.id,
	   panel_components.name,
	   panel_measurement_units.name AS unit,
	   min_amount,
	   max_amount,
	   comment
FROM panel_components
JOIN panel_measurement_units ON panel_components.unit_id = panel_measurement_units.id;
